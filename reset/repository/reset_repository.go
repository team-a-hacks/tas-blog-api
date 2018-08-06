package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	model "github.com/team-a-hacks/tas-blog-api/reset"
	"github.com/team-a-hacks/tas-blog-api/status"
)

type resetRepository struct {
	Conn *gorm.DB
}

// NewResetRepository mount reset repository
func NewResetRepository(db *gorm.DB) ResetRepository {
	return &resetRepository{
		Conn: db,
	}
}

// ResetRepository repository interface
type ResetRepository interface {
	GetByAccountID(id uuid.UUID) (*model.Reset, error)
	GetByToken(token string) (*model.Reset, error)
	Create(reset model.Reset) error
	Update(reset model.Reset) error
	Delete(id uuid.UUID) error
}

func (m *resetRepository) GetByAccountID(id uuid.UUID) (*model.Reset, error) {
	reset := model.Reset{}
	err := m.Conn.Model(&reset).Where("account_id = ?", id).Find(&reset).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &reset, nil
}

func (m *resetRepository) GetByToken(token string) (*model.Reset, error) {
	reset := model.Reset{}
	err := m.Conn.Model(&reset).Where("token = ?", token).Find(&reset).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, errors.New("token is not found").Error())
	} else if err != nil {
		return nil, err
	}
	return &reset, nil
}

func (m *resetRepository) Create(reset model.Reset) error {
	err := m.Conn.Create(&reset).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *resetRepository) Update(reset model.Reset) error {
	err := m.Conn.Model(&reset).Update(&reset).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *resetRepository) Delete(id uuid.UUID) error {
	reset := model.Reset{}
	err := m.Conn.Model(&reset).Where("id = ?", id).Delete(&reset).Error
	if err != nil {
		return err
	}
	return nil
}
