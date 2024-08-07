package go_web_sdk

import (
	"fmt"
	"github.com/Amosawy/go_web_sdk/response"
	"github.com/Amosawy/go_web_sdk/router"
	"github.com/gin-gonic/gin"
	"testing"
)

func test1(c *gin.Context) {
	c.JSON(200, response.BuildSuccessResp("i am amos"))
}

func TestWeb(t *testing.T) {
	engine := CreateAmosGin()
	err := router.RegisterRouter(engine, "/api/*action")
	engine.GET("/test", test1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = Run(engine, 8080)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
