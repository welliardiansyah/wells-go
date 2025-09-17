package security

import "time"

type Maker interface {
	CreateToken(userID, email string, roles []string, permissions []Permission, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
