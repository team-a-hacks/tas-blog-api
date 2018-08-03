package status

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo"
)

var (
	// ErrInternalServer internal server error
	ErrInternalServer = errors.New("Internal server error")
	// ErrNotFound not found error
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict conflit error
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadRequest bad request error
	ErrBadRequest = errors.New("Bad request")
	// ErrUnauthorized unauthorized error
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrForbidden forbidden error
	ErrForbidden = errors.New("Forbidden")
)

// ResponseError 返却するエラーコードの指定
func ResponseError(ctx echo.Context, err error) error {
	switch errors.Cause(err) {
	case ErrInternalServer:
		return ctx.JSON(500, err.Error())
	case ErrNotFound:
		return ctx.JSON(http.StatusNotFound, err.Error())
	case ErrConflict:
		return ctx.JSON(http.StatusConflict, err.Error())
	case ErrBadRequest:
		return ctx.JSON(http.StatusBadRequest, err.Error())
	case ErrUnauthorized:
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	case ErrForbidden:
		return ctx.JSON(http.StatusForbidden, err.Error())
	default:
		return ctx.JSON(500, err.Error())
	}
}
