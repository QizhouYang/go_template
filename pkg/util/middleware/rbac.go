package middleware

import (
	"go_template/pkg/constant"
	"go_template/pkg/dto"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/storyicon/grbac"
)

func RBACMiddleware() iris.Handler {
	r, err := grbac.New(grbac.WithAdvancedRules(constant.Roles))
	if err != nil {
		panic(err)
	}
	return func(c *context.Context) {
		roles := querySystemRoles(c)
		state, err := r.IsRequestGranted(c.Request(), roles)
		if err != nil {
			c.StatusCode(http.StatusInternalServerError)
			c.StopExecution()
			return
		}
		if !state.IsGranted() {
			c.StatusCode(http.StatusForbidden)
			c.StopExecution()
			return
		}
		c.Next()
	}
}

func querySystemRoles(ctx *context.Context) []string {
	u := ctx.Values().Get("user").(dto.SessionUser)
	ctx.Values().Set("roles", u.Roles)
	return u.Roles
}
