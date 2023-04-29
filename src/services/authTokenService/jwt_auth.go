package authtokenservice

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// AuthTokenProvider is a JWT implementation of the AuthTokenRepository interface.
type authTokenProvider struct {
	signingMethod        jwt.SigningMethod
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

// NewAuthTokenProvider creates a new AuthTokenProvider.
func NewAuthTokenProvider(signingMethod string) AuthTokenRepository {
	return &authTokenProvider{
		signingMethod:        jwt.SigningMethodHS256,
		accessTokenLifetime:  15 * 60,          // 15 minutes
		refreshTokenLifetime: 24 * 7 * 60 * 60, // 24 hours
	}
}

// getJWTSecret is a utility functions that returns the JWT secret used to sign the tokens.
func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// mapToClaims is a utility function that converts a map to a jwt.MapClaims object.
func mapToClaims(payload map[string]interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	for key, value := range payload {
		claims[key] = value
	}
	return claims
}

// CreateAccessToken creates an access token with the given payload.
func (a *authTokenProvider) CreateAccessToken(payload map[string]interface{}) (string, error) {
	key := getJWTSecret()

	claims := mapToClaims(payload)
	claims["exp"] = time.Now().Add(a.accessTokenLifetime * time.Second).Unix()
	claims["is_access_token"] = true

	token := jwt.NewWithClaims(a.signingMethod, claims)
	return token.SignedString(key)
}

// CreateRefreshToken creates a refresh token with the given payload.
func (a *authTokenProvider) CreateRefreshToken(payload map[string]interface{}) (string, error) {
	key := getJWTSecret()

	claims := mapToClaims(payload)
	claims["exp"] = time.Now().Add(a.refreshTokenLifetime * time.Second).Unix()
	claims["is_refresh_token"] = true

	token := jwt.NewWithClaims(a.signingMethod, claims)
	return token.SignedString(key)
}

// ParseAccessToken parses the given access token and returns the payload.
func (a *authTokenProvider) ParseAccessToken(token string) (map[string]interface{}, error) {
	claims, err := parseToken(token, a.signingMethod)
	if err != nil {
		return nil, err
	}

	if claims["is_access_token"] != true {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// ParseRefreshToken parses the given refresh token and returns the payload.
func (a *authTokenProvider) ParseRefreshToken(token string) (map[string]interface{}, error) {
	claims, err := parseToken(token, a.signingMethod)
	if err != nil {
		return nil, err
	}

	if claims["is_refresh_token"] != true {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// parseToken is a utility function that parses the given token and returns the payload.
func parseToken(token string, signingMethod jwt.SigningMethod) (map[string]interface{}, error) {
	key := getJWTSecret()

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
