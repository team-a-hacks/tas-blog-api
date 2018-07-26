// @APIVersion 1.0.0
// @APITitle tas-blog-api
// @APIDescription tas-blog-api
// @BasePath http://localhost:1323/swagger-ui
// @SubApi hello [/hello]
package main

import (
	"os"

	"github.com/labstack/echo"
	helloC "github.com/team-a-hacks/tas-blog-api/hello/controller"
	swaggerui "github.com/team-a-hacks/tas-blog-api/swagger-ui"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	setup(e)

	helloC.NewHelloController(e)

	swaggerui.NewSwaggerUIController(e)

	// サーバー起動
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
