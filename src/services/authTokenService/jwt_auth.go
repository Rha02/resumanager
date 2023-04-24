package authtokenservice

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthTokenProvider struct {
	signingMethod        jwt.SigningMethod
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func NewAuthTokenProvider(signingMethod string) AuthTokenRepository {
	return &AuthTokenProvider{
		signingMethod:        jwt.SigningMethodHS256,
		accessTokenLifetime:  15 * 60,          // 15 minutes
		refreshTokenLifetime: 24 * 7 * 60 * 60, // 24 hours
	}
}

// getJWTSecret returns the JWT secret used to sign the tokens.
func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func mapToClaims(payload map[string]interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	for key, value := range payload {
		claims[key] = value
	}
	return claims
}

func (a *AuthTokenProvider) CreateAccessToken(payload map[string]interface{}) (string, error) {
	key := getJWTSecret()

	claims := mapToClaims(payload)
	claims["exp"] = time.Now().Add(a.accessTokenLifetime * time.Second).Unix()

	token := jwt.NewWithClaims(a.signingMethod, claims)
	return token.SignedString(key)
}

func (a *AuthTokenProvider) CreateRefreshToken(payload map[string]interface{}) (string, error) {
	key := getJWTSecret()

	claims := mapToClaims(payload)
	claims["exp"] = time.Now().Add(a.refreshTokenLifetime * time.Second).Unix()

	token := jwt.NewWithClaims(a.signingMethod, claims)
	return token.SignedString(key)
}

func (a *AuthTokenProvider) ParseToken(token string) (map[string]interface{}, error) {
	key := getJWTSecret()

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != a.signingMethod {
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
