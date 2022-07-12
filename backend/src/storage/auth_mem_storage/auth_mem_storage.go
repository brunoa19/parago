package auth_mem_storage

import (
	"shipa-gen/src/service/authorisation_service"
)

type AuthMemStorage struct {
	sessions map[string]*authorisation_service.AuthUserInfo
}

func NewAuthMemStorage() authorisation_service.AuthorizationStorage {
	return &AuthMemStorage{
		sessions: make(map[string]*authorisation_service.AuthUserInfo),
	}
}

func (s *AuthMemStorage) Store(sessionToken string, userInfo authorisation_service.AuthUserInfo) error {
	s.sessions[sessionToken] = &userInfo

	return nil
}

func (s *AuthMemStorage) Load(sessionToken string) (*authorisation_service.AuthUserInfo, error) {
	userInfo, _ := s.sessions[sessionToken]

	return userInfo, nil
}

func (s *AuthMemStorage) Remove(sessionToken string) error {
	delete(s.sessions, sessionToken)

	return nil
}
