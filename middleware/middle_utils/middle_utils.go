package middle_utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const REQUEST_SUCCESS = 0

const (
	URL_METRICS    = "/metrics"
	URL_HEART_BEAT = "/heartbeat"
	URL_CLB_HEART  = "/"
)

var IgnorePaths []string

func init() {
	IgnorePaths = []string{
		URL_HEART_BEAT,
		URL_METRICS,
		URL_CLB_HEART,
	}
}

func SetRespGetterFactory(respGetter RespGetterFactory) {
	respGetterFactory = respGetter
}

func GetRespGetterFactory() RespGetterFactory {
	return respGetterFactory
}

type RespGetterFactory func() RespGetter

var respGetterFactory RespGetterFactory

type RespGetter interface {
	GetCode() int
}

type BodyLogWriter struct {
	gin.ResponseWriter
	BodyBuf *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	//copy to BodyBuf 多复制一次到buf中供后续拿到response
	return w.BodyBuf.Write(b)
	//return w.ResponseWriter.Write(b)
}

var (
	TotalCounterVec = prometheus.NewCounterVec( //监控连接个数
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Total number of HTTP requests made",
		},
		[]string{"path"},
	)
	ReqLoginErrorVec = prometheus.NewCounterVec( //返回错误个数
		prometheus.CounterOpts{
			Name: "request_error_count",
			Help: "Total request error count of the host",
		},
		[]string{"path", "code"},
	)
	ReqDurationVec = prometheus.NewHistogramVec( //响应时间
		prometheus.HistogramOpts{
			Name: "request_latency",
			Help: "record request latency",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(
		TotalCounterVec,
		ReqDurationVec,
		ReqLoginErrorVec,
	)
}
