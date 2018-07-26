package db

import (
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DB データベース接続インスタンス
	DB *gorm.DB
)

// ConnectDB データベース接続
func ConnectDB() *gorm.DB {
	url := os.Getenv("DATABASE_URL") + os.Getenv("SSL_MODE")
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	dbLogMode, _ := strconv.ParseBool(os.Getenv("DB_LOGMODE"))
	db.LogMode(dbLogMode)

	return db
}
