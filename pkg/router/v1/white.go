package v1

import (
	"go_template/pkg/util/captcha"
	"net/http"

	"github.com/kataras/iris/v12/context"
)

func downloadFile(ctx *context.Context) {
	str := ""
	_, _ = ctx.WriteString(str)
}

func generateCaptcha(ctx *context.Context) {
	c, err := captcha.CreateCaptcha()
	if err != nil {
		_, _ = ctx.JSON(err)
		ctx.StatusCode(http.StatusInternalServerError)
	}
	_, _ = ctx.JSON(&c)

}
