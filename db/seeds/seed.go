package main

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/team-a-hacks/tas-blog-api/db"
	"github.com/team-a-hacks/tas-blog-api/db/migrations"
	"github.com/team-a-hacks/tas-blog-api/db/seeds/sql"
)

var odb *gorm.DB

func main() {
	odb := db.ConnectDB()
	defer odb.Close()

	migrations.DropTable(odb)
	migrations.Migrate(odb)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		odb.DB().Exec(sql.InsertAccounts)
		wg.Done()
	}()
	wg.Wait()
}
