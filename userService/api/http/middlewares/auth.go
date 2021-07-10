package middlewares

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
	"user-service/services"
)

func AuthMW(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil || accessToken == "" {
		redirectUrl := services.OAuth2().AuthCodeURL("ok")
		s := sessions.Default(ctx)
		s.Set("user_service_redirect_url", ctx.Request.RequestURI)
		err = s.Save()
		if err != nil {
			hapi.JsonResponseErr(ctx,err)
			ctx.Abort()
			return
		}
		ctx.Redirect(http.StatusFound, redirectUrl)
		ctx.Abort()
		return
	}
	ctx.Next()
}
