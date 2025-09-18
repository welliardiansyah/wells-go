package security

import "time"

type Permission struct {
	Name      string `json:"name"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
	CanExport bool   `json:"can_export"`
	CanImport bool   `json:"can_import"`
	CanView   bool   `json:"can_view"`
}

type Payload struct {
	UserID      string       `json:"sub"`
	Email       string       `json:"email"`
	Roles       []string     `json:"roles"`
	Permissions []Permission `json:"permissions"`
	IssuedAt    time.Time    `json:"issued_at"`
	ExpiredAt   time.Time    `json:"expired_at"`
}

const AuthorizationPayloadKey = "auth_payload"
