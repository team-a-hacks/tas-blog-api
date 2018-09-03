package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/team-a-hacks/tas-blog-api/comment"
	"github.com/team-a-hacks/tas-blog-api/comment/usecase"
	"github.com/team-a-hacks/tas-blog-api/status"
)

// CommentController comment controller
type CommentController struct {
	CommentUsecase usecase.CommentUsecase
}

// NewCommentController mounnt commnet controller
func NewCommentController(e *echo.Echo, comment usecase.CommentUsecase) {
	handler := &CommentController{
		CommentUsecase: comment,
	}
	e.POST("/comments", handler.Create)
}

// Create comment create
// @Title コメント作成API
// @Description コメント作成
// @Assept json
// @Param comment body comment.Payload true "コメント"
// @Success 200 {object} comment.ResponseData true "コメント"
// @Failure 404 {object} error true "not found error"
// @Failure 500 {object} error true "internal server error"
// @Resource /comments
// @Router /comments [post]
func (c *CommentController) Create(ctx echo.Context) error {
	req := comment.Payload{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res, err := c.CommentUsecase.Create(&req)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusCreated, res)
}
