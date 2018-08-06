// @APIVersion 1.0.0
// @APITitle tas-blog-api
// @APIDescription tas-blog-api
// @BasePath http://localhost:1323/swagger-ui
// @SubApi hello [/hello]
// @SubApi auth [/auth]
package main

import (
	"os"

	"github.com/labstack/echo"
	accountR "github.com/team-a-hacks/tas-blog-api/account/repository"
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

	helloC.NewHelloController(e)
	authC.NewAuthController(e, authUsecase)

	swaggerui.NewSwaggerUIController(e)

	// サーバー起動
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
