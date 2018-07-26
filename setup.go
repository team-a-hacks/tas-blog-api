package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/team-a-hacks/tas-blog-api/db"
	"github.com/team-a-hacks/tas-blog-api/db/migrations"
)

var database *gorm.DB

func setup(e *echo.Echo) {
	// DB接続
	database = db.ConnectDB()
	// マイグレーション
	migrations.Excute()
}
