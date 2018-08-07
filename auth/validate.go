package auth

import (
	"errors"
	"net/mail"
	"regexp"

	uuid "github.com/satori/go.uuid"
)

// LoginValidate validate login
func LoginValidate(login *Login) error {
	// メールアドレスのバリデーション
	if login.Email == "" {
		return errors.New("email is required")
	}
	if login.Email != "" {
		_, err := mail.ParseAddress(login.Email)
		if err != nil {
			return errors.New("mail format error")
		}
	}

	// パスワードのバリデーション
	if login.Password == "" {
		return errors.New("password is required")
	}
	if login.Password != "" {
		if len([]byte(login.Password)) < 8 {
			return errors.New("password must be 8 length")
		}
		// 半角英数字かのバリデーション
		var passRegexp = regexp.MustCompile(`^[A-Za-z0-9]*$`)
		ok := passRegexp.MatchString(login.Password)
		if !ok {
			return errors.New("invalid password")
		}
	}
	return nil
}

// LogoutValidate validate logout
func LogoutValidate(logout *Logout) error {
	if logout.ID == uuid.FromStringOrNil("") {
		return errors.New("id is required")
	}
	return nil
}
