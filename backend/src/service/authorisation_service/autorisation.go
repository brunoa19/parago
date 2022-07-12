package authorisation_service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"shipa-gen/src/client/shipa_client"
	"shipa-gen/src/models"
	"shipa-gen/src/service"
	"time"
)

const AuthStorageMemory = "memory"
const AuthStorageMongo = "mongo"

type AuthUserInfo struct {
	ExternalToken string          `bson:"externalToken"`
	Login         string          `bson:"login"`
	PasswordHash  string          `bson:"passwordHash"`
	TimeLogin     time.Time       `bson:"timeLogin"`
	TimeExtValid  time.Time       `bson:"timeExtValid"`
	UserInfo      models.UserInfo `bson:"userInfo"`
}

func (u *AuthUserInfo) IsOwnedBy(login, passwordHash string) bool {
	return u.Login == login && u.PasswordHash == passwordHash
}

func (u *AuthUserInfo) IsRequireExtValidation(time time.Time) bool {
	return u.TimeExtValid.Unix() < time.Unix()
}

type authService struct {
	shipaClient           shipa_client.ShipaClient
	storage               AuthorizationStorage
	sessions              map[string]*AuthUserInfo
	extTokenCheckInterval time.Duration
}

func NewService(shipaClient shipa_client.ShipaClient, storage AuthorizationStorage, extTokenCheckInterval time.Duration) service.AuthorisationService {
	return &authService{
		shipaClient:           shipaClient,
		storage:               storage,
		extTokenCheckInterval: extTokenCheckInterval,
		sessions:              make(map[string]*AuthUserInfo),
	}
}

func (s *authService) VerifyToken(token string) (*shipa_client.ShipaUserInfo, error) {
	// verify token start
	extUserInfo, err := s.shipaClient.GetUser(token)
	if err != nil {
		return nil, err
	}
	if extUserInfo == nil {
		return nil, errors.New("invalid token")
	}
	if extUserInfo.Deactivated {
		return nil, errors.New("account deactivated")
	}

	return extUserInfo, nil
}

func (s *authService) AuthLogin(login string, password string) (*models.UserInfo, error) {
	now := time.Now()
	// check if exists
	passwordHash := s.generatePasswordHash(login, password)
	for _, userInfo := range s.sessions {
		if userInfo.IsOwnedBy(login, passwordHash) && !userInfo.IsRequireExtValidation(now) {
			return &userInfo.UserInfo, nil
		}
	}

	extTokenInfo, err := s.shipaClient.GetToken(login, password)
	if err != nil {
		return nil, err
	}
	if extTokenInfo == nil {
		log.Println("empty GetToken response")
		return nil, nil
	}

	extUserInfo, err := s.VerifyToken(extTokenInfo.Token)
	if err != nil {
		return nil, err
	}

	sessionToken := s.generateSessionToken(extTokenInfo.Token)

	userInfo := &AuthUserInfo{
		ExternalToken: extTokenInfo.Token,
		Login:         login,
		PasswordHash:  passwordHash,
		TimeLogin:     now,
		TimeExtValid:  now.Add(s.extTokenCheckInterval),
		UserInfo: models.UserInfo{
			Org:    extUserInfo.OrgName,
			User:   extUserInfo.Name,
			UserId: extUserInfo.Id,
			Token:  sessionToken,
		},
	}

	if err = s.storage.Store(sessionToken, *userInfo); err != nil {
		return nil, err
	}

	return &userInfo.UserInfo, nil
}

func (s *authService) AuthLogout(sessionToken string) error {
	userInfo, err := s.storage.Load(sessionToken)
	if err != nil {
		return err
	}
	if userInfo == nil {
		return errors.New("invalid token")
	}

	return s.storage.Remove(sessionToken)
}

func (s *authService) AuthGetUser(sessionToken string) (*models.UserInfo, string, error) {
	userInfo, err := s.storage.Load(sessionToken)
	if err != nil {
		return nil, "", err
	}
	if userInfo == nil {
		return nil, "", nil
	}

	now := time.Now()
	if userInfo.IsRequireExtValidation(now) {

		extUserInfo, err := s.VerifyToken(userInfo.ExternalToken)
		if err != nil {
			return nil, "", err
		}

		if extUserInfo == nil {
			_ = s.storage.Remove(sessionToken)
		}

		userInfo.TimeExtValid = now.Add(s.extTokenCheckInterval)

		_ = s.storage.Store(sessionToken, *userInfo) //
	}

	return &userInfo.UserInfo, userInfo.ExternalToken, nil
}

func (s *authService) generateSessionToken(externalToken string) string {
	return fmt.Sprintf("%x%d", md5.Sum([]byte(externalToken)), len(externalToken))
}

func (s *authService) generatePasswordHash(login, password string) string {
	tmp := login + password
	return fmt.Sprintf("%x%d", md5.Sum([]byte(tmp)), len(tmp))
}
