package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
	"user-service/api/http/middlewares"
	"user-service/services"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type UserServiceController struct {
}

func (s UserServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s UserServiceController) RouterGroupName() (name string) {
	return "user_service"
}

func (s UserServiceController) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		sessions.Sessions("user_service", services.SessionStore()),
		middlewares.AuthMW,
	}
}

func (s UserServiceController) Version() string {
	return "v1"
}

// Ping godoc
// @Summary Ping
// @Description Ping
// @ID user-service-ping
// @Tags 默认
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=string} ""
// @Router /v1/user_service/ping [get]
func (s UserServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(ctx *gin.Context) {
		//ctx.SetCookie("access_token", "nekilc", 3000, "/", "", false, true)
		hapi.JsonResponseOk(ctx, "pong", hapi.WithMsg("ok"))
	}
}

var userInfoMap = map[string]interface{}{
	"id":          "800001",
	"username":    "nieaowei",
	"nickname":    "Nekilc",
	"avatar":      "",
	"create_time": "2021-06-01",
}

// GetUserinfo godoc
// @Summary fetch user info
// @Description fetch user info
// @ID user-service-GetUserinfo
// @Tags 默认
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=string} ""
// @Router /v1/user_service/userinfo [get]
func (s UserServiceController) GetUserinfo() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "userinfo", "v1", func(ctx *gin.Context) {
		hapi.JsonResponseOk(ctx, userInfoMap)
	}
}

func (s UserServiceController) PostUserinfo() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "userinfo", "v1", func(ctx *gin.Context) {

	}
}

func (s UserServiceController) PutUserinfo() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPut, "userinfo", "v1", func(ctx *gin.Context) {
		hapi.JsonResponseOk(ctx, "ok")
	}
}

func (s UserServiceController) PatchUserinfo() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPatch, "userinfo", "v1", func(ctx *gin.Context) {
		hapi.JsonResponseOk(ctx, "ok")
	}
}

type UnAuthController struct {
}

func (u UnAuthController) RouterRegister(group *gin.RouterGroup) {
}

func (u UnAuthController) RouterGroupName() (name string) {
	return "user_service/unauth"
}

func (u UnAuthController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (u UnAuthController) Version() string {
	return ""
}

// UnauthPing godoc
// @Summary UnauthPing
// @Description UnauthPing
// @ID user-service-UnauthPing
// @Tags 默认
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=string} ""
// @Router /v1/user_service/unauth/ping [get]
func (s UnAuthController) UnauthPing() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(ctx *gin.Context) {
		ctx.SetCookie("access_token", "nekilc", 3000, "/", "", false, true)
		hapi.JsonResponseOk(ctx, "ok")
	}
}
