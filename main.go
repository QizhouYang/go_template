package main

import (
	_ "go_template/docs"
	"go_template/pkg/server"
	"log"
)

//go:generate go-bindata -o ./pkg/util/i18n/locales.go -pkg i18n ./locales/...
//go:generate swag init

/**
func main() {
	logger.Init()
	app := iris.New()
	config.Init()
	database.CreateInitDB().Init()
	app.Get("/ping", pong).Describe("healthcheck")

	mvc.Configure(app.Party("/greet"), setup)

	// http://localhost:8080/greet?name=kataras
	app.Listen(":8080", iris.WithLogLevel("debug"))
}

func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}

func setup(app *mvc.Application) {
	// Register Dependencies.
	app.Register(
		environment.DEV, // DEV, PROD
		//database.NewDB,          // sqlite, mysql
		service.NewGreetService, // greeterWithLogging, greeter
	)

	// Register Controllers.
	app.Handle(new(controller.GreetController))
}
**/

// @title       GoTemplate Restful API
// @version  1.0
// @termsOfService
// @contact.name                Fit2cloud Support
// @contact.url                 https://www.fit2cloud.com
// @contact.email               test@fit2cloud.com
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @schemes http
// @description This is a sample server
// @BasePath                    /api/v1
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
