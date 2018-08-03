package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	model "github.com/team-a-hacks/tas-blog-api/account"
	"github.com/team-a-hacks/tas-blog-api/status"
)

type accountRepository struct {
	Conn *gorm.DB
}

// NewAccountRepository mount account repository
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

// AccountRepository repository interface
type AccountRepository interface {
	GetByEmail(email string) (*model.Account, error)
	GetByID(id uuid.UUID) (*model.Account, error)
	Create(u *model.Account) (*model.Account, error)
	Update(u *model.Account) error
}

func (m *accountRepository) GetByEmail(email string) (*model.Account, error) {
	var a model.Account
	err := m.Conn.Model(&a).Where("email = ?", email).Find(&a).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &a, nil
}

func (m *accountRepository) List() ([]*model.Account, error) {
	var account = []*model.Account{}
	err := m.Conn.Model(&account).Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (m *accountRepository) GetByID(id uuid.UUID) (*model.Account, error) {
	var a model.Account
	err := m.Conn.Model(&a).Where("id = ?", id).Find(&a).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &a, nil
}

func (m *accountRepository) Create(u *model.Account) (*model.Account, error) {
	err := m.Conn.Create(u).Error
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return u, nil
}

func (m *accountRepository) Update(u *model.Account) error {
	err := m.Conn.Model(u).Update(u).Error
	if err != nil {
		return err
	}
	return nil
}
