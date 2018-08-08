package usecase

import (
	uuid "github.com/satori/go.uuid"
	model "github.com/team-a-hacks/tas-blog-api/account"
	accountRepo "github.com/team-a-hacks/tas-blog-api/account/repository"
)

type accountUsecase struct {
	accountRepo accountRepo.AccountRepository
}

// NewAccountUsecase mount account usecase
func NewAccountUsecase(account accountRepo.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: account,
	}
}

// AccountUsecase account usecase interface
type AccountUsecase interface {
	Get(id uuid.UUID) (*model.AccountData, error)
}

func (a *accountUsecase) Get(id uuid.UUID) (*model.AccountData, error) {
	account, err := a.accountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	accountData := formatAccountData(account)
	return accountData, nil
}

func formatAccountData(account *model.Account) *model.AccountData {
	var accountData = model.AccountData{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}
	return &accountData
}
