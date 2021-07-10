package controllers

import (
	"crypto/tls"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
	"user-service/services"
)

type AuthController struct {
}

func (a AuthController) RouterRegister(group *gin.RouterGroup) {
}

func (a AuthController) RouterGroupName() (name string) {
	return "user_service"
}

func (a AuthController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{
		sessions.Sessions("user_service", services.SessionStore()),
	}
}

func (a AuthController) Version() string {
	return ""
}

type GetAuthParams struct {
	Code  string `form:"code" binding:"required"`
	State string `form:"state" `
}

type GetAuthResp struct {
	Token string `json:"token"`
}

func (a AuthController) Auth() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "oauth2", "v1", func(ctx *gin.Context) {
		var (
			err    error
			params = new(GetAuthParams)
			resp   = new(GetAuthResp)
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(ctx, err)
				return
			}
			hapi.JsonResponseOk(ctx, resp)
		}()
		if err = ctx.ShouldBind(params); err != nil {
			return
		}
		http.DefaultTransport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		token, err := services.OAuth2().Exchange(ctx, params.Code)
		if err != nil {
			return
		}
		session := sessions.Default(ctx)
		redirectUrl, ok := session.Get("user_service_redirect_url").(string)
		if !ok {
			err = hapi.NewCodeErr(1, "")
			return
		}
		ctx.SetCookie("access_token", token.AccessToken, 0, "", "", false, true)
		ctx.Redirect(302, redirectUrl)
	}
}
