package controller

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/team-a-hacks/tas-blog-api/article"
	"github.com/team-a-hacks/tas-blog-api/article/usecase"
	"github.com/team-a-hacks/tas-blog-api/status"
)

// ArticleController article struct
type ArticleController struct {
	ArticleUsecase usecase.ArticleUsecase
}

// NewArticleController Article controller
func NewArticleController(e *echo.Echo, us usecase.ArticleUsecase) {
	handler := &ArticleController{
		ArticleUsecase: us,
	}
	e.GET("/articles", handler.List)
}

// List get article list
// @Title 記事一覧取得API
// @Description 記事一覧取得API
// @Assept json
// @Success 200 {array} article.ArticleData true "記事""
// @Failure 404 {object} error true "not found error"
// @Failure 500 {object} error true "internal server error"
// @Resource /articles
// @Router /articles [get]
func (c *ArticleController) List(ctx echo.Context) error {
	res, err := c.ArticleUsecase.List()
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
