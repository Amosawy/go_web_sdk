package router

import (
	"github.com/Amosawy/go_web_sdk/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

var RouteConf map[string]Route

func RegisterRouter(engine *gin.Engine, url string) error {
	engine.Any(url, Forward)
	routeConf, err := LoadRouteConfig()
	if err != nil {
		seelog.Errorf("LoadRouteConfig error %s", err.Error())
		return err
	}
	RouteConf = routeConf
	return nil
}

func Forward(c *gin.Context) {
	action, _ := c.Params.Get("action")
	route, ok := RouteConf[action]
	if !ok {
		c.String(http.StatusBadGateway, "502 BadGateway route not found")
		return
	}
	hostReverseProxy(&route, c.Writer, c.Request)
}

func hostReverseProxy(route *Route, w http.ResponseWriter, req *http.Request) { //反向代理
	seelog.Infof("redirect route : %+v\n", route)
	direct := func(req *http.Request) {
		req.URL.Path = route.Uri
		req.URL.Scheme = route.Scheme
		req.URL.Host = route.Host
		seelog.Infof("redirect url : %+v\n", req.URL)
	}
	proxy := &httputil.ReverseProxy{Director: direct}
	proxy.ServeHTTP(w, req)
}