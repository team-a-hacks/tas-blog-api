package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/team-a-hacks/tas-blog-api/refreshtoken"
	"github.com/team-a-hacks/tas-blog-api/status"
)

type refreshTokenRepository struct {
	Conn *gorm.DB
}

// NewRefreshTokenRepository mount refresh token repository
func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{
		Conn: db,
	}
}

// RefreshTokenRepository repository interface
type RefreshTokenRepository interface {
	GetByAccountID(id uuid.UUID) (*refreshtoken.RefreshToken, error)
	Create(token *refreshtoken.RefreshToken) error
	Delete(token string) error
	Get(token string) (*refreshtoken.RefreshToken, error)
}

func (m *refreshTokenRepository) GetByAccountID(id uuid.UUID) (*refreshtoken.RefreshToken, error) {
	var token refreshtoken.RefreshToken
	err := m.Conn.Model(&token).Where("account_id = ?", id).Find(&token).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &token, nil
}

func (m *refreshTokenRepository) Create(token *refreshtoken.RefreshToken) error {
	err := m.Conn.Create(token).Error
	if err != nil {
		log.Println(err)
		return errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return nil
}

func (m *refreshTokenRepository) Delete(token string) error {
	var rToken refreshtoken.RefreshToken
	err := m.Conn.Model(&rToken).Where("refresh_token = ?", token).Delete(&rToken).Error
	if err != nil {
		log.Println(err)
		return errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return nil
}

func (m *refreshTokenRepository) Get(token string) (*refreshtoken.RefreshToken, error) {
	var refreshToken refreshtoken.RefreshToken
	err := m.Conn.Model(&refreshToken).Where("refresh_token = ?", token).Find(&refreshToken).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &refreshToken, nil
}
