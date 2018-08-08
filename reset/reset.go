package reset

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Reset reset password token struct
type Reset struct {
	ID        uuid.UUID `form:"id" json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	AccountID uuid.UUID `form:"account_id" json:"accountID" sql:"type:uuid" name:"account ID"`
	Token     string    `form:"token" json:"token" name:"token"`
	CreatedBy uuid.UUID `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time `json:"updatedAt" name:"更新日"`
}
