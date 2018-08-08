package middleware

import (
	"crypto/rsa"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/team-a-hacks/tas-blog-api/utils"
)

type jwtMiddleware struct {
	key *rsa.PublicKey
}

// NewJWTMiddleware mount jwt middleware
func NewJWTMiddleware(token utils.TokenHandler) (JWTMiddleware, error) {
	key, err := token.LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return &jwtMiddleware{key: key}, nil
}

// JWTMiddleware jwt middleware interface
type JWTMiddleware interface {
	JWT() echo.MiddlewareFunc
}

func (j *jwtMiddleware) JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    j.key,
		SigningMethod: "RS256",
	})
}
