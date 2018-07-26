package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloController hello controller
type HelloController struct{}

// NewHelloController mount hello controller
func NewHelloController(e *echo.Echo) {
	handler := &HelloController{}

	e.GET("/hello", handler.Hello)
}

// Hello hello world
// Hello hello world
// @Title Hello
// @Description Hello 'your name'
// @Assept json
// @Success 200 {object} string true "Hello World"
// @Resource /hello
// @Router /hello [get]
func (c *HelloController) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
