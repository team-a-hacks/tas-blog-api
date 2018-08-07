package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/team-a-hacks/tas-blog-api/auth"
	"github.com/team-a-hacks/tas-blog-api/auth/usecase"
	"github.com/team-a-hacks/tas-blog-api/status"
)

// AuthController auth controller
type AuthController struct {
	AuthUsecase usecase.AuthUsecase
}

// NewAuthController mount auth contorller
func NewAuthController(e *echo.Echo, us usecase.AuthUsecase) {
	handler := &AuthController{
		AuthUsecase: us,
	}

	e.POST("/auth/login", handler.Login)
	e.POST("/auth/refresh", handler.Refresh)
	e.POST("/auth/logout", handler.Logout)
}

// Login login
// @Title ログインAPI
// @Description ログイン
// @Assept json
// @param auth body auth.Login true "auth"
// @Success 200 {object} auth.Token "トークン"
// @Failure 400 {object} error "bad request"
// @Resource /auth
// @Router /auth/login [post]
func (c *AuthController) Login(ctx echo.Context) error {
	request := auth.Login{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// validation
	err = auth.LoginValidate(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := c.AuthUsecase.Login(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, token)
}

// Refresh refresh
// @title レフレッシュAPI
// @Description リフレッシュトークンの再発行
// @Assept json
// @Param token body auth.Refresh true "リフレッシュトークン"
// @Success 200 {object} auth.Token "トークン"
// @Failure 400 {object} error "bad request"
// @Resource /auth
// @Router /auth/refresh [post]
func (c *AuthController) Refresh(ctx echo.Context) error {
	request := auth.Refresh{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	token, err := c.AuthUsecase.Refresh(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, token)
}

// Logout logout
// @title ログアウトAPI
// @Description 店舗従業員ログアウト
// @Assept json
// @Param logout body auth.Logout true "アカウントID"
// @Success 204 {object} error "no content"
// @Failure 400 {object} error "bad request"
// @Resource /auth
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx echo.Context) error {
	request := auth.Logout{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// validation
	err = auth.LogoutValidate(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.AuthUsecase.Logout(request.ID)
	if err != nil {
		return status.ResponseError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
