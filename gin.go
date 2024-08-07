package go_web_sdk

import (
	"fmt"
	"github.com/Amosawy/go_web_sdk/middleware/middle_utils"
	"github.com/Amosawy/go_web_sdk/middleware/mlog"
	"github.com/Amosawy/go_web_sdk/seelog"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"syscall"
)

func CreateAmosGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	engine.Use(mlog.InfoLog())
	engine.GET(middle_utils.URL_METRICS, gin.WrapH(promhttp.Handler()))
	engine.GET(middle_utils.URL_HEART_BEAT, HeartBeat)
	return engine
}

func Run(engine *gin.Engine, port int) error {
	server := endless.NewServer(fmt.Sprintf(":%d", port), engine)
	server.BeforeBegin = func(add string) {
		seelog.Infof("Actual pid is %d", syscall.Getpid())
	}
	return server.ListenAndServe()
}

func HeartBeat(c *gin.Context) {
	c.String(http.StatusOK, "SUCCESS")
}
