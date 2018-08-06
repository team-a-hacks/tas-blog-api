package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/team-a-hacks/tas-blog-api/db"
	"github.com/team-a-hacks/tas-blog-api/db/migrations"
	mid "github.com/team-a-hacks/tas-blog-api/middleware"
	"github.com/team-a-hacks/tas-blog-api/utils"
)

var database *gorm.DB
var jwt mid.JWTMiddleware
var token utils.TokenHandler

func setup(e *echo.Echo) {
	// DB接続
	database = db.ConnectDB()
	// マイグレーション
	migrations.Excute()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))

	var err error
	// mount token handler
	token = utils.NewTokenHandler()

	// mount jwt middleware
	jwt, err = mid.NewJWTMiddleware(token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
		os.Exit(1)
	}

}
