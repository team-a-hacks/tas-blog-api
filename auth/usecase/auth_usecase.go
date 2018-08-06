package usecase

import (
	"os"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	accountRepo "github.com/team-a-hacks/tas-blog-api/account/repository"
	"github.com/team-a-hacks/tas-blog-api/auth"
	"github.com/team-a-hacks/tas-blog-api/refreshtoken"
	rTokenRepo "github.com/team-a-hacks/tas-blog-api/refreshtoken/repository"
	"github.com/team-a-hacks/tas-blog-api/status"
	"github.com/team-a-hacks/tas-blog-api/utils"
)

type authUsecase struct {
	accountRepo accountRepo.AccountRepository
	rTokenRepo  rTokenRepo.RefreshTokenRepository
	token       utils.TokenHandler
}

// NewAuthUsecase mount auth usecase
func NewAuthUsecase(
	account accountRepo.AccountRepository,
	rToken rTokenRepo.RefreshTokenRepository,
	token utils.TokenHandler,
) AuthUsecase {
	return &authUsecase{
		accountRepo: account,
		rTokenRepo:  rToken,
		token:       token,
	}
}

// AuthUsecase auth usecase interface
type AuthUsecase interface {
	Login(l *auth.Login) (*auth.Token, error)
}

func (a *authUsecase) Login(l *auth.Login) (*auth.Token, error) {
	account, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CompareHashPassword(account.Password, l.Password)
	if err != nil {
		return nil, status.ErrUnauthorized
	}

	refresh, err := a.rTokenRepo.GetByAccountID(account.ID)
	// リフレッシュトークンがすでにあれば削除する
	if refresh != nil {
		err := a.rTokenRepo.Delete(refresh.RefreshToken)
		if err != nil {
			return nil, err
		}
	}

	refreshToken := a.token.RandToken()
	// 数値化
	atoi, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES_MIN"))
	insert := &refreshtoken.RefreshToken{
		ID:           uuid.NewV4(),
		AccountID:    account.ID,
		RefreshToken: refreshToken,
		Expired:      time.Now().Add(time.Minute * time.Duration(atoi)),
	}
	err = a.rTokenRepo.Create(insert)
	if err != nil {
		return nil, err
	}

	token, err := a.token.GenerateJWT(account.ID, false)
	if err != nil {
		return nil, err
	}

	res := auth.Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return &res, nil
}
