package authtokenservice

type AuthTokenRepository interface {
	// CreateAccessToken creates an access token with the given payload.
	CreateAccessToken(payload map[string]interface{}) (string, error)

	// CreateRefreshToken creates a refresh token with the given payload.
	CreateRefreshToken(payload map[string]interface{}) (string, error)

	// ParseToken parses the given access token and returns the payload.
	ParseAccessToken(token string) (map[string]interface{}, error)

	// ParseToken parses the given refresh token and returns the payload.
	ParseRefreshToken(token string) (map[string]interface{}, error)
}
