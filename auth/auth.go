package auth

import uuid "github.com/satori/go.uuid"

// Login login struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token token struct
type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Refresh refresh struct
type Refresh struct {
	AccountID    uuid.UUID `json:"accountID"`
	RefreshToken string    `json:"refreshToken"`
}
