package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/team-a-hacks/tas-blog-api/article"
	"github.com/team-a-hacks/tas-blog-api/db"
)

// Excute マイグレーション実行
func Excute() {
	odb := db.ConnectDB()
	Migrate(odb)
	odb.Close()
}

// Migrate マイグレーション
func Migrate(odb *gorm.DB) {
	odb.AutoMigrate(
		// ここにテーブルを追加する
		article.Article{},
	)
}

// DropTable テーブル削除
func DropTable(odb *gorm.DB) {
	odb.DropTableIfExists(
		article.Article{},
	)
}
