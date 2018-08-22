package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/team-a-hacks/tas-blog-api/account"
	"github.com/team-a-hacks/tas-blog-api/article"
	"github.com/team-a-hacks/tas-blog-api/comment"
	"github.com/team-a-hacks/tas-blog-api/db"
	"github.com/team-a-hacks/tas-blog-api/refreshtoken"
	"github.com/team-a-hacks/tas-blog-api/reset"
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
		account.Account{},
		comment.Comment{},
		refreshtoken.RefreshToken{},
		reset.Reset{},
	)
}

// DropTable テーブル削除
func DropTable(odb *gorm.DB) {
	odb.DropTableIfExists(
		article.Article{},
		account.Account{},
		comment.Comment{},
		refreshtoken.RefreshToken{},
		reset.Reset{},
	)
}
