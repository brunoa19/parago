package shipa_client

type ShipaClient interface {
	GetToken(email, password string) (*ShipaTokenResponse, error)
	GetUser(token string) (*ShipaUserInfo, error)

	GetResponse(token string, path string) (int, string, []byte, error)
}
