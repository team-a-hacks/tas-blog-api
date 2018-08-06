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
