package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Manager struct {
	signingKey string
}

var (
	ErrEmptySignKey           = errors.New("empty signing key")
	ErrGetUserClaimsFromToken = errors.New("error get user claims from token")
	ErrAssertSubID            = errors.New("error to assert sub_id")
)

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, ErrEmptySignKey
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userID int, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub_id": userID,
		"exp":    jwt.NewNumericDate(time.Now().Add(ttl)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.signingKey)
}

func (m *Manager) NewRefreshToken(userID int, ttl time.Duration) (string, error) {
	uuid := uuid.New().String()

	claims := jwt.MapClaims{
		"jti":    uuid,
		"sub_id": userID,
		"exp":    jwt.NewNumericDate(time.Now().Add(ttl)),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := refreshToken.SignedString(m.signingKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (m *Manager) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i any, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrGetUserClaimsFromToken
	}

	sub, ok := claims["sub"].(int)
	if !ok {
		return 0, ErrAssertSubID
	}

	return sub, nil
}
