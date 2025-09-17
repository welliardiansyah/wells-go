package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
	issuer    string
}

func NewJWTMaker(secretKey, issuer string) (*JWTMaker, error) {
	if len(secretKey) < 32 {
		return nil, errors.New("secret key must be at least 32 characters")
	}
	return &JWTMaker{secretKey: secretKey, issuer: issuer}, nil
}

func (j *JWTMaker) CreateToken(userID, email string, roles []string, permissions []Permission, duration time.Duration) (string, error) {
	if len(roles) == 0 {
		return "", errors.New("roles cannot be empty")
	}

	now := time.Now().UTC()
	exp := now.Add(duration)

	claims := &Claims{
		Sub:         userID,
		Email:       email,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) VerifyToken(tokenStr string) (*Payload, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, errors.New("invalid token")
	}

	return &Payload{
		UserID:      claims.Sub,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
		IssuedAt:    claims.IssuedAt.Time,
		ExpiredAt:   claims.ExpiresAt.Time,
	}, nil
}
