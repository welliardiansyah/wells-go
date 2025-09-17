package security

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Sub         string       `json:"sub"`
	Email       string       `json:"email"`
	Roles       []string     `json:"roles"`
	Permissions []Permission `json:"permissions"`
	jwt.RegisteredClaims
}
