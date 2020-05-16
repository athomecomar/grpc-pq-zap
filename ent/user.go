package ent

type User struct {
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}
