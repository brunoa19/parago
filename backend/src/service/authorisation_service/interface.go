package authorisation_service

type AuthorizationStorage interface {
	Store(sessionToken string, userInfo AuthUserInfo) error
	Load(sessionToken string) (*AuthUserInfo, error)
	Remove(sessionToken string) error
}
