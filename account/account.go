package account

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account account struct
type Account struct {
	ID        uuid.UUID  `form:"id" json:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	Name      string     `form:"name" json:"name" name:"名前"`
	Email     string     `form:"email" json:"email" gorm:"unique" name:"メールアドレス"`
	Password  string     `form:"password" json:"password" name:"パスワード"`
	CreatedBy uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time `json:"deletedAt" name:"削除日"`
}

// AccountData account data struct
type AccountData struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
