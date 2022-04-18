package proxy

import (
	"crypto/tls"
	"fmt"
	"go_template/pkg/constant"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kataras/iris/v12/context"
)

func LoggingProxy(ctx *context.Context) {
	clusterName := ctx.Params().Get("cluster_name")
	proxyPath := ctx.Params().Get("p")
	if clusterName == "" {
		_, _ = ctx.JSON(http.StatusBadRequest)
		return
	}
	endpoint := "127.0.0.1"
	/**if err != nil {
		_, _ = ctx.JSON(http.StatusInternalServerError)
		return
	}**/
	u, err := url.Parse(fmt.Sprintf("http://%s", endpoint))
	if err != nil {
		_, _ = ctx.JSON(http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req := ctx.Request()
	req.Host = fmt.Sprintf(constant.DefaultLoggingIngress)
	req.URL.Path = proxyPath
	proxy.ServeHTTP(ctx.ResponseWriter(), req)
}
