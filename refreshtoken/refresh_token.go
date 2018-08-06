package refreshtoken

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// RefreshToken refresh token struct
type RefreshToken struct {
	ID           uuid.UUID `form:"id" json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	AccountID    uuid.UUID `sql:"type:uuid"`
	RefreshToken string    `json:"refreshToken"`
	Expired      time.Time
	CreatedAt    time.Time `json:"created_at" name:"作成日"`
}
