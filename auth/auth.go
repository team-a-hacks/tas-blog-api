package auth

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
