package seelog

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/go-basic/uuid"
	"github.com/petermattis/goid"
	"github.com/quanhengzhuang/requestid"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	seelog.RegisterCustomFormatter("ServiceName", createAppNameFormatter)
	logger, _ := seelog.LoggerFromConfigAsString(seelogConfig)
	seelog.ReplaceLogger(logger)
}

const (
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_ERROR = "ERROR"
	LOG_LEVEL_DEBUG = "DEBUG"
	LOG_LEVEL_WARN  = "WARN"
)

func Infof(format string, params ...interface{}) {
	seelog.Infof(getPrefix(LOG_LEVEL_INFO)+format, params...)
}

func Info(params ...interface{}) {
	prefix := getPrefix(LOG_LEVEL_INFO)
	var newParams []interface{}
	newParams = append(newParams, prefix)
	for _, v := range params {
		newParams = append(newParams, v)
	}
	seelog.Info(newParams...)
}


func Errorf(format string, params ...interface{}) {
	seelog.Infof(getPrefix(LOG_LEVEL_ERROR)+format, params...)
}

func Error(params ...interface{}) {
	prefix := getPrefix(LOG_LEVEL_ERROR)
	var newParams []interface{}
	newParams = append(newParams, prefix)
	for _, v := range params {
		newParams = append(newParams, v)
	}
	seelog.Info(newParams...)
}

func Flush() {
	seelog.Flush()
}

func getPrefix(level string) string { //获取日志前缀普遍信息
	callerInfo := getCallerName()
	requestId := requestid.Get()
	if requestId == nil {
		requestIdStr := fmt.Sprintf("%+v", uuid.New())
		requestid.Set(requestIdStr)
	}
	prefix := fmt.Sprintf("%s:::%v:::%d:::%s:::", level, requestId, goid.Get(), callerInfo)
	return prefix
}

func getCallerName() string { //获取调用函数堆栈信息
	pc, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf("%s.%d %s", filepath.Base(file), line, filepath.Base(runtime.FuncForPC(pc).Name()))
}

func createAppNameFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		serviceName := os.Getenv("SERVICE_NAME")
		if serviceName == "" {
			serviceName = "None"
		}
		return serviceName
	}
}

var seelogConfig string = `
<seelog minlevel="trace">
	<outputs formatid="fmt_info">
         <filter levels="trace,debug,info,warn,error,critical">
			 <rollingfile formatid="fmt_info" type="size" filename="../log/web.log"  maxsize="104857600" maxrolls="10"/>
         </filter>
         <filter levels="error,critical">
			 <rollingfile formatid="fmt_err" type="size" filename="../log/error/web_error.log"  maxsize="10485760" maxrolls="100"/>
         </filter>
	</outputs>
	<formats>
		<format id="fmt_info" format="%Date(2006-01-02 15:04:05.999):::%Msg%n" />
		<format id="fmt_err" format="%Date(2006-01-02 15:04:05.999):::%Msg%n" />
	</formats>
</seelog>`
