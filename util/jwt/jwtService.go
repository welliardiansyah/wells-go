package jwt

import (
	"errors"
	"time"
	"wells-go/infrastructure/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func cfg() *config.Config {
	return config.GetConfig()
}

func AccessTTL() time.Duration {
	d, err := time.ParseDuration(cfg().AccessTokenTTL)
	if err != nil {
		d = 15 * time.Minute
	}
	return d
}

func RefreshTTL() time.Duration {
	d, err := time.ParseDuration(cfg().RefreshTokenTTL)
	if err != nil {
		d = 30 * 24 * time.Hour
	}
	return d
}

func secret() []byte { return []byte(cfg().JWTSecret) }
func issuer() string { return cfg().JWTIssuer }

func GenerateAccessToken(userID, email, role string) (string, int64, error) {
	now := time.Now().UTC()
	exp := now.Add(AccessTTL())

	c := Claims{
		Sub:   userID,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(secret())
	if err != nil {
		return "", 0, err
	}
	return s, exp.Unix(), nil
}

func GenerateRefreshToken(userID string) (string, int64, error) {
	now := time.Now().UTC()
	exp := now.Add(RefreshTTL())

	c := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    issuer(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(exp),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(secret())
	if err != nil {
		return "", 0, err
	}
	return s, exp.Unix(), nil
}

func Parse(tokenStr string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret(), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
