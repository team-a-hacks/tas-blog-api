package comment

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Comment comment struct
type Comment struct {
	ID        uuid.UUID      `form:"id" json:"id" gorm:"primary_key" sql:"type:uuid" name:"id"`
	Content   string         `form:"content" json:"content" name:"内容"`
	Author    sql.NullString `form:"author" json:"author" name:"コメント作成者"`
	AuthorID  uuid.UUID      `form:"authorId" json:"authorId" sql:"type:uuid" name:"作成者ID"`
	ArticleID uuid.UUID      `form:"articleId" json:"articleId" sql:"type:uuid" name:"記事ID"`
	CreatedBy uuid.UUID      `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID      `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time      `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time      `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time     `json:"deletedAt" name:"削除日"`
}

// Payload comment payload struct
type Payload struct {
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	AuthorID  uuid.UUID `json:"authorId"`
	ArticleID uuid.UUID `json:"articleId"`
	CreatedBy uuid.UUID `json:"createdBy"`
}

// ResponseData response data
type ResponseData struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	ArticleID uuid.UUID `json:"articleId"`
	PostDate  string    `json:"postDate"`
}
