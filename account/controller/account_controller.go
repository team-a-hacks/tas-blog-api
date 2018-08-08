package controller

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	_ "github.com/team-a-hacks/tas-blog-api/account"
	"github.com/team-a-hacks/tas-blog-api/account/usecase"
	"github.com/team-a-hacks/tas-blog-api/status"
)

// AccountController account controller
type AccountController struct {
	AccountUsecase usecase.AccountUsecase
}

// NewAccountController mount account controller
func NewAccountController(e *echo.Echo, account usecase.AccountUsecase) {
	handler := &AccountController{
		AccountUsecase: account,
	}

	e.GET("/accounts/:id", handler.Get)
}

// Get get account controller
// @Title アカウント取得API
// @Param id path uuid.UUID true "ID"
// @Success 200 {object} account.AccountData true "アカウント情報"
// @Failure 404 {object} error true "not found error"
// @Failure 500 {object} error true "internal server error"
// @Resource /accounts
// @Router /accounts/{id} [get]
func (c *AccountController) Get(ctx echo.Context) error {
	id := uuid.FromStringOrNil(ctx.Param("id"))
	res, err := c.AccountUsecase.Get(id)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
