package controller

import (
	"go_template/pkg/environment"
	"go_template/pkg/model"
	"go_template/pkg/service"
	"go_template/pkg/service/impl"
	"go_template/pkg/util/logger"

	"github.com/kataras/iris/v12"
)

type GreetController struct {
	Service impl.GreetService
	Ctx     *iris.Context
}

func NewGreetController() *GreetController {
	return &GreetController{
		Service: service.NewGreetService(environment.DEV),
	}
}

// Get all
// @Tags all
// @Summary Show a test
// @Description show a test id
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Security ApiKeyAuth
// @Router /greet/ [get]
func (c *GreetController) Get() (model.Response, error) {
	logger.Log.Infof("exec")

	logger.Log.Infof("succes end")
	return model.Response{Message: "测试 Get"}, nil
}

// Get name
// @Tags byname
// @Summary Show a test
// @Description show a test id
// @Accept  json
// @Produce  json
// @Param username query string false "username param" minlength(2) maxlength(10)
// @Success 200 {object} model.Response
// @Security ApiKeyAuth
// @Router /greet/{username}/ [get]
func (c *GreetController) GetBy(username string) (model.Response, error) {
	logger.Log.Infof("exec")
	message, err := c.Service.Say(username)

	print(message)
	if err != nil {
		logger.Log.Errorf(err.Error())
		return model.Response{}, err
	}

	logger.Log.Infof("succes end")
	return model.Response{Message: message}, nil
}

func (c *GreetController) Delete(req model.Request) (model.Response, error) {

	err := c.Service.Delete(req.Name)
	return model.Response{Message: "Delete success!"}, err
}
