package v1

import (
	"encoding/json"
	"go_template/pkg/util/middleware"
	"net/http"

	"go_template/pkg/controller"
	"go_template/pkg/util/errorf"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/pkg/errors"
)

var AuthScope iris.Party
var WhiteScope iris.Party

func V1(parent iris.Party) {
	v1 := parent.Party("/v1")
	authParty := v1.Party("/auth")
	mvc.New(authParty.Party("/session")).HandleError(ErrorHandler).Handle(controller.NewSessionController())
	mvc.New(v1.Party("/user")).HandleError(ErrorHandler).Handle(controller.NewForgotPasswordController())
	mvc.New(v1.Party("/greet")).HandleError(ErrorHandler).Handle(controller.NewGreetController())
	AuthScope = v1.Party("/")
	AuthScope.Use(middleware.JWTMiddleware().Serve)
	AuthScope.Use(middleware.UserMiddleware)
	AuthScope.Use(middleware.RBACMiddleware())
	AuthScope.Use(middleware.PagerMiddleware)
	AuthScope.Use(middleware.ForceMiddleware)
	//mvc.New(AuthScope.Party("/projects/{project}/clusters/{cluster}/resources")).HandleError(ErrorHandler).Handle(controller.NewClusterResourceController())
	WhiteScope = v1.Party("/")
	WhiteScope.Get("/downlaod/static/{name}", downloadFile)
	WhiteScope.Get("/captcha", generateCaptcha)

}

func ErrorHandler(ctx *context.Context, err error) {
	if err != nil {
		warp := struct {
			Msg string `json:"msg"`
		}{err.Error()}
		var result string
		switch errType := err.(type) {
		case gorm.Errors:
			errorSet := make(map[string]string)
			for _, er := range errType {
				tr := ctx.Tr(er.Error())
				if tr != "" {
					errorMsg := tr
					errorSet[er.Error()] = errorMsg
				}
			}
			for _, set := range errorSet {
				result = result + set + " "
			}
		case error:
			switch errRoot := errors.Cause(err).(type) {
			case errorf.CErrFs:
				errs := errRoot.Get()
				for _, er := range errs {
					args := er.Args.([]interface{})
					tr := ctx.Tr(er.Msg, args...)
					if tr != "" {
						result = result + tr + "\n "
					}
				}
			default:
				tr := ctx.Tr(errors.Cause(err).Error())
				if tr != "" {
					result = tr
				} else {
					result = err.Error()
				}
			}
		}
		warp.Msg = result
		bf, _ := json.Marshal(&warp)
		ctx.StatusCode(http.StatusBadRequest)
		_, _ = ctx.WriteString(string(bf))
		ctx.StopExecution()
		return
	}
}
