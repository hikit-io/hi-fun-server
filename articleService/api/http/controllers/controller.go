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

type ArticleServiceController struct {
}

func (s ArticleServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s ArticleServiceController) RouterGroupName() (name string) {
	return "article_service"
}

func (s ArticleServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s ArticleServiceController) Version() string {
	return "v1"
}

// Ping godoc
// @Summary Ping service
// @Description Ping service
// @ID article_service-ping
// @Tags 默认
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=string} ""
// @Router /v1/article_service/ping [get]
func (s ArticleServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(ctx *gin.Context) {
		hapi.JsonResponseOk(ctx, "pong", hapi.WithMsg("pong"))
	}
}
