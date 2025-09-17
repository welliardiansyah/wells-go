package security

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type JWTMaker struct {
	secretKey string
	issuer    string
	redis     *redis.Client
	ctx       context.Context
}

func NewJWTMaker(secretKey, issuer string, redisClient *redis.Client) (*JWTMaker, error) {
	if len(secretKey) < 32 {
		return nil, errors.New("secret key must be at least 32 characters")
	}
	return &JWTMaker{
		secretKey: secretKey,
		issuer:    issuer,
		redis:     redisClient,
		ctx:       context.Background(),
	}, nil
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
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	err = j.redis.Set(j.ctx, "jwt:"+signedToken, "active", duration).Err()
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTMaker) VerifyToken(tokenStr string) (*Payload, error) {
	val, err := j.redis.Get(j.ctx, "jwt:"+tokenStr).Result()
	if err == redis.Nil {
		return nil, errors.New("token revoked or not found in cache")
	} else if err != nil {
		return nil, err
	}
	if val != "active" {
		return nil, errors.New("token invalid")
	}

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
		Email:       claims.Email,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
		IssuedAt:    claims.IssuedAt.Time,
		ExpiredAt:   claims.ExpiresAt.Time,
	}, nil
}

func (j *JWTMaker) RevokeToken(tokenStr string) error {
	return j.redis.Del(j.ctx, "jwt:"+tokenStr).Err()
}
