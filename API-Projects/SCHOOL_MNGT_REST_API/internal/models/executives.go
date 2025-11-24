package models

import "database/sql"

type Executive struct {
	ID                  int            `json:"id,omitempty"`
	FirstName           string         `json:"first_name,omitempty"`
	LastName            string         `json:"last_name,omitempty"`
	Email               string         `json:"email,omitempty"`
	Username            string         `json:"username,omitempty"`
	Password            string         `json:"password,omitempty"`
	PasswordChangedAt   sql.NullString `json:"password_changed_at,omitempty"`
	UserCreatedAt       sql.NullString `json:"user_created_at,omitempty"`
	PasswordResetCode   sql.NullString `json:"password_reset_token,omitempty"`
	PasswordCodeExpires sql.NullString `json:"password_reset_expires,omitempty"`
	InactiveStatus      bool           `json:"inactive_status,omitempty"`
	Role                string         `json:"role,omitempty"`
}
