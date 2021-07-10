package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type TempServiceController struct {
}

func (s TempServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s TempServiceController) RouterGroupName() (name string) {
	return "temp_service"
}

func (s TempServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s TempServiceController) Version() string {
	return "v1"
}

// Ping godoc
// @Summary Ping service
// @Description Ping service
// @ID temp_service-ping
// @Tags 默认
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=string} ""
// @Router /v1/temp_service/ping [get]
func (s TempServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(ctx *gin.Context) {
		hapi.JsonResponseOk(ctx, "pong", hapi.WithMsg("pong"))
	}
}
