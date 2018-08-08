// @APIVersion 1.0.0
// @APITitle tas-blog-api
// @APIDescription tas-blog-api
// @BasePath http://localhost:1323/swagger-ui
// @SubApi hello [/hello]
// @SubApi auth [/auth]
// @SubApi accounts [/accounts]
package main

import (
	"os"

	"github.com/labstack/echo"
	accountC "github.com/team-a-hacks/tas-blog-api/account/controller"
	accountR "github.com/team-a-hacks/tas-blog-api/account/repository"
	accountU "github.com/team-a-hacks/tas-blog-api/account/usecase"
	authC "github.com/team-a-hacks/tas-blog-api/auth/controller"
	authU "github.com/team-a-hacks/tas-blog-api/auth/usecase"
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

	authUsecase := authU.NewAuthUsecase(accountRepo, rTokenRepo, token)
	accountUcase := accountU.NewAccountUsecase(accountRepo)

	helloC.NewHelloController(e)
	authC.NewAuthController(e, authUsecase)
	accountC.NewAccountController(e, accountUcase)

	swaggerui.NewSwaggerUIController(e)

	// サーバー起動
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
