package middleware

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CORSMiddleware init cors config
func CORSMiddleware() middleware.CORSConfig {
	origin := os.Getenv("ORIGIN")
	config := middleware.CORSConfig{
		AllowOrigins: strings.Split(origin, "|"),
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}
	return config
}
