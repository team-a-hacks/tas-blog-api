// @APIVersion 1.0.0
// @APITitle echo-swagger
// @APIDescription echo-swagger sample api
// @BasePath http://localhost:1323/swagger-ui
// @SubApi hello [/hello]
package main

import (
	"github.com/labstack/echo"
	helloC "github.com/team-a-hacks/tas-blog-api/hello/controller"
	swaggerui "github.com/team-a-hacks/tas-blog-api/swagger-ui"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	helloC.NewHelloController(e)

	swaggerui.NewSwaggerUIController(e)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323")) //ポート番号指定してね
}
