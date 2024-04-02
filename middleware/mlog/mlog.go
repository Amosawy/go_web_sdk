package mlog

import (
	"bytes"
	"fmt"
	"github.com/Amosawy/go_web_sdk/middleware/middle_utils"
	"github.com/Amosawy/go_web_sdk/seelog"
	"github.com/Amosawy/go_web_sdk/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/quanhengzhuang/requestid"
	"io/ioutil"
	"strings"
	"time"
)

const MAX_PRINT_BODY_LEN = 1024

const (
	REQUEST_ID = "Request-Id"
)

func InfoLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		if tools.InList(context.Request.URL.Path, middle_utils.IgnorePaths) {
			return
		}
		beginTime := time.Now()
		// 1. get request body
		body, _ := ioutil.ReadAll(context.Request.Body)
		context.Request.Body.Close()
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		//2. set requestID for goroutine ctx
		requestId := context.Request.Header.Get(REQUEST_ID)
		middle_utils.TotalCounterVec.WithLabelValues(context.Request.URL.Path).Inc()
		if requestId == "" {
			requestId = fmt.Sprintf("%+v", uuid.New())
		}
		requestid.Set(requestId)
		defer requestid.Delete()
		seelog.Infof("Req Url: %s %+v,Body %s  Header: %s", context.Request.Method, context.Request.URL, string(body), tools.GetFmtStr(context.Request.Header))
		//3. set response writer
		blw := middle_utils.BodyLogWriter{
			ResponseWriter: context.Writer,
			BodyBuf:        bytes.NewBufferString(""),
		}
		context.Writer = blw
		//4. do next
		context.Next()
		//5. log resp body
		strBody := strings.Trim(blw.BodyBuf.String(), "\n")
		if len(strBody) > MAX_PRINT_BODY_LEN {
			strBody = strBody[:(MAX_PRINT_BODY_LEN - 1)]
		}
		//6. judge logic error
		//getterFactory := middle_utils.GetRespGetterFactory()
		//rspGetter := getterFactory()
		//json.Unmarshal(blw.BodyBuf.Bytes(), &rspGetter)
		//if rspGetter.GetCode() != middle_utils.REQUEST_SUCCESS {
		//	middle_utils.ReqLoginErrorVec.WithLabelValues(context.Request.URL.Path, fmt.Sprintf("%d", 1)).Inc()
		//}
		seelog.Infof("url: %+v, cost %v Resp Body %s", context.Request.URL, time.Since(beginTime), strBody)
		duration := float64(time.Since(beginTime)) / float64(time.Second)
		middle_utils.ReqDurationVec.WithLabelValues(context.Request.URL.Path).Observe(duration)
	}
}
