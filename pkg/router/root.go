package router

import (
	"fmt"
	"go_template/pkg/controller"
	"go_template/pkg/router/proxy"
	v1 "go_template/pkg/router/v1"
	"go_template/pkg/util/i18n"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func Server() *iris.Application {
	app := iris.New()
	err := app.I18n.LoadAssets(i18n.AssetNames, i18n.Asset, "en-US", "zh-CN")
	if err != nil {
		fmt.Println(err.Error())
	}
	app.I18n.SetDefault("zh-CN")
	app.I18n.ExtractFunc = func(ctx iris.Context) string {
		return ctx.GetHeader("lang")
	}
	c := &swagger.Config{
		URL: "/swagger/doc.json",
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(c, swaggerFiles.Handler))
	app.Get("/api/v1/health", controller.HealthController)
	proxy.RegisterProxy(app)
	api := app.Party("/api")
	v1.V1(api)
	return app
}
