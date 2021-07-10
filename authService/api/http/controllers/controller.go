package controllers

import (
	"auth-service/common"
	"auth-service/logs"
	"auth-service/services"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type AuthServiceController struct {
}

func (s AuthServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s AuthServiceController) RouterGroupName() (name string) {
	return "auth_service"
}

func (s AuthServiceController) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		sessions.Sessions("auth_service", services.SessionStore()),
	}
}

func (s AuthServiceController) Version() string {
	return "v1"
}

func (s AuthServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}

func (s AuthServiceController) LoginGet() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "login", "v1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}

type LoginPostParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s AuthServiceController) LoginPost() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "login", "v1", func(ctx *gin.Context) {
		var (
			params = new(LoginPostParams)
			err    error
		)
		defer func() {
			if err != nil {
				logs.TraceLog(ctx, err.Error())
				hapi.JsonResponseErr(ctx, err)
				return
			}
			ctx.Redirect(http.StatusFound, "auth")
		}()
		if err = ctx.ShouldBind(params); err != nil {
			return
		}
		//todo 验证账号
		if params.Username == params.Password {
			return
		}
		err = hapi.NewCodeErr(200, "err")
	}
}

func (s AuthServiceController) GetAuth() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "auth", "v1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "auth.html", nil)
	}
}

type GetAuthorizeParams struct {
	ClientId            string `form:"client_id"`
	CodeChallenge       string `form:"code_challenge"`
	CodeChallengeMethod string `form:"code_challenge_method"`
	RedirectUri         string `form:"redirect_uri"`
}

func (s AuthServiceController) GetAuthorize() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "authorize", "v1", func(ctx *gin.Context) {
		var (
			err    error
			params = new(GetAuthorizeParams)
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(ctx, err)
				return
			}
		}()
		if err = ctx.ShouldBind(params); err != nil {
			return
		}
		sess := sessions.Default(ctx)
		sess.Set("auth_service_redirect_uri", params.RedirectUri)
		err = sess.Save()
		if err != nil {
			return
		}
		ctx.Redirect(http.StatusFound, "login")
	}
}

func (s AuthServiceController) PostAuthorize() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "authorize", "v1", func(ctx *gin.Context) {
		var (
			err error
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(ctx, err)
				return
			}
		}()
		sess := sessions.Default(ctx)
		url, ok := sess.Get("auth_service_redirect_uri").(string)
		if !ok || url == "" {
			err = errors.New("authorize fail")
			return
		}
		ctx.Redirect(http.StatusFound, url+"?code=1234")
	}
}

type PostTokenParams struct {
	State string `form:"state"`
	Code  string `form:"code"`
}

func (s AuthServiceController) PostToken() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "token", "v1", func(ctx *gin.Context) {
		var (
			err    error
			params = &PostTokenParams{}
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(ctx, err)
				return
			}
		}()
		if err = ctx.ShouldBind(params); err != nil {
			return
		}
		t, err := common.JwtEncode(jwt.StandardClaims{
			Audience:  "auth_service",
			ExpiresAt: 0,
			Id:        "1",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		})
		if err != nil {
			return
		}
		tt := oauth2.Token{
			AccessToken:  t,
			TokenType:    "",
			RefreshToken: "",
			Expiry:       time.Time{},
		}
		ctx.JSON(200, tt)
		//hapi.JsonResponseOk(ctx, tt)
	}
}
