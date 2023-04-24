package authtokenservice

import "errors"

type TestAuthToken struct{}

func NewTestAuthTokenRepo() AuthTokenRepository {
	return &TestAuthToken{}
}

func (t *TestAuthToken) CreateAccessToken(payload map[string]interface{}) (string, error) {
	if payload["username"] == "access_token_error" {
		return "", errors.New("error")
	}
	return "access_token", nil
}

func (t *TestAuthToken) CreateRefreshToken(payload map[string]interface{}) (string, error) {
	if payload["username"] == "refresh_token_error" {
		return "", errors.New("error")
	}
	return "refresh_token", nil
}

func (t *TestAuthToken) ParseToken(token string) (map[string]interface{}, error) {
	if token == "error" {
		return nil, errors.New("error")
	}

	return map[string]interface{}{
		"test": "1",
	}, nil
}
