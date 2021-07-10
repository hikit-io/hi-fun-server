package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PracticeServiceController struct {
}

func (s PracticeServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s PracticeServiceController) RouterGroupName() (name string) {
	return "practice-service"
}

func (s PracticeServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s PracticeServiceController) Version() string {
	return "v1"
}

func (s PracticeServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}
