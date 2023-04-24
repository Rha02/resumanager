package authtokenservice

type AuthTokenRepository interface {
	// CreateAccessToken creates an access token with the given payload.
	CreateAccessToken(payload map[string]interface{}) (string, error)

	// CreateRefreshToken creates a refresh token with the given payload.
	CreateRefreshToken(payload map[string]interface{}) (string, error)

	// ParseToken parses the given token, validates it, and returns the payload.
	ParseToken(token string) (map[string]interface{}, error)
}

var Repo AuthTokenRepository

func NewAuthTokenRepo(repo AuthTokenRepository) {
	Repo = repo
}
