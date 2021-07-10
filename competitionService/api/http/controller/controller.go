package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompetitionServiceController struct {
}

func (s CompetitionServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s CompetitionServiceController) RouterGroupName() (name string) {
	return "competition-service"
}

func (s CompetitionServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s CompetitionServiceController) Version() string {
	return "v1"
}

func (s CompetitionServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}
