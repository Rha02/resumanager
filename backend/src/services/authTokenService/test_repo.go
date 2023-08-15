package authtokenservice

import "errors"

type testAuthToken struct{}

func NewTestAuthTokenRepo() AuthTokenRepository {
	return &testAuthToken{}
}

func (t *testAuthToken) CreateAccessToken(payload map[string]interface{}) (string, error) {
	if payload["username"] == "access_token_error" {
		return "", errors.New("error")
	}
	return "access_token", nil
}

func (t *testAuthToken) CreateRefreshToken(payload map[string]interface{}) (string, error) {
	if payload["username"] == "refresh_token_error" {
		return "", errors.New("error")
	}
	return "refresh_token", nil
}

func (t *testAuthToken) ParseAccessToken(token string) (map[string]interface{}, error) {
	if token == "error" {
		return nil, errors.New("error")
	}

	if token == "refresh_token" {
		return nil, errors.New("not access token")
	}

	return map[string]interface{}{
		"username": "testuser",
		"password": "testpassword",
	}, nil
}

func (t *testAuthToken) ParseRefreshToken(token string) (map[string]interface{}, error) {
	if token == "error" {
		return nil, errors.New("error")
	}

	if token == "access_token" {
		return nil, errors.New("not refresh token")
	}

	if token == "creating_access_token_error" {
		return map[string]interface{}{
			"username": "access_token_error",
		}, nil
	}

	return map[string]interface{}{
		"username": "testuser",
		"password": "testpassword",
	}, nil
}
