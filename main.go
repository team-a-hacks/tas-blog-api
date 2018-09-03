// @APIVersion 1.0.0
// @APITitle tas-blog-api
// @APIDescription tas-blog-api
// @BasePath http://localhost:1323/swagger-ui
// @SubApi hello [/hello]
// @SubApi auth [/auth]
// @SubApi accounts [/accounts]
// @SubApi articles [/articles]
// @SubApi comments [/comments]
package main

import (
	"os"

	"github.com/labstack/echo"
	accountC "github.com/team-a-hacks/tas-blog-api/account/controller"
	accountR "github.com/team-a-hacks/tas-blog-api/account/repository"
	accountU "github.com/team-a-hacks/tas-blog-api/account/usecase"
	articleC "github.com/team-a-hacks/tas-blog-api/article/controller"
	articleR "github.com/team-a-hacks/tas-blog-api/article/repository"
	articleU "github.com/team-a-hacks/tas-blog-api/article/usecase"
	authC "github.com/team-a-hacks/tas-blog-api/auth/controller"
	authU "github.com/team-a-hacks/tas-blog-api/auth/usecase"
	commentC "github.com/team-a-hacks/tas-blog-api/comment/controller"
	commentR "github.com/team-a-hacks/tas-blog-api/comment/repository"
	commentU "github.com/team-a-hacks/tas-blog-api/comment/usecase"
	helloC "github.com/team-a-hacks/tas-blog-api/hello/controller"
	rTokenR "github.com/team-a-hacks/tas-blog-api/refreshtoken/repository"
	swaggerui "github.com/team-a-hacks/tas-blog-api/swagger-ui"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	setup(e)

	accountRepo := accountR.NewAccountRepository(database)
	rTokenRepo := rTokenR.NewRefreshTokenRepository(database)
	articleRepo := articleR.NewArticleRepository(database)
	commentRepo := commentR.NewCommentRepository(database)

	authUsecase := authU.NewAuthUsecase(accountRepo, rTokenRepo, token)
	accountUcase := accountU.NewAccountUsecase(accountRepo)
	articleUcase := articleU.NewArticleUsecase(articleRepo)
	commentUcase := commentU.NewCommentUsecase(commentRepo)

	helloC.NewHelloController(e)
	authC.NewAuthController(e, authUsecase)
	accountC.NewAccountController(e, accountUcase)
	articleC.NewArticleController(e, articleUcase)
	commentC.NewCommentController(e, commentUcase)

	swaggerui.NewSwaggerUIController(e)

	// サーバー起動
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
