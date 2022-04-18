package proxy

import "github.com/kataras/iris/v12"

var (
	keyPrefix           = "Bearer"
	AuthorizationHeader = "Authorization"
	//clusterService      = service.NewClusterService()
)

func RegisterProxy(parent iris.Party) {
	proxy := parent.Party("/proxy")
	proxy.Any("/logging/{cluster_name}/{p:path}", LoggingProxy)
}
