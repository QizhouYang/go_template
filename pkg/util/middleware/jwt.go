package middleware

import (
	"net/http"

	"go_template/pkg/constant"
	"go_template/pkg/model"

	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/viper"
)

var (
	UserIsNotRelatedProject = "USER_IS_NOT_RELATED_PROJECT"
)

type JwtMiddleware struct {
	*jwtmiddleware.Middleware
}

func (m *JwtMiddleware) Serve(ctx *context.Context) {
	session := constant.Sess.Start(ctx)
	u := session.Get(constant.SessionUserKey)
	if u != nil {
		ctx.Next()
		return
	}
	if err := m.CheckJWT(ctx); err != nil {
		m.Config.ErrorHandler(ctx, err)
		return
	}
	ctx.Next()
}

func JWTMiddleware() *JwtMiddleware {
	secretKey := []byte(viper.GetString("jwt.secret"))
	m := JwtMiddleware{jwtmiddleware.New(
		jwtmiddleware.Config{
			Extractor: jwtmiddleware.FromAuthHeader,
			ValidationKeyGetter: func(token *jwtmiddleware.Token) (interface{}, error) {
				return secretKey, nil
			},
			SigningMethod: jwtmiddleware.SigningMethodHS256,
			ErrorHandler:  ErrorHandler,
		},
	)}
	return &m
}

func ErrorHandler(ctx *context.Context, err error) {
	if err == nil {
		return
	}
	ctx.StopExecution()
	response := &model.Response{
		Message: err.Error(),
	}
	ctx.StatusCode(http.StatusUnauthorized)
	_, _ = ctx.JSON(response)
}
