package article

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Article article struct
type Article struct {
	ID        uuid.UUID  `form:"id" json:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	Title     string     `form:"title" json:"title" name:"title"`
	Content   string     `form:"content" json:"content" name:"記事"`
	CreatedBy uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time `json:"deletedAt" name:"削除日"`
}
