package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
	"user-service/api/http/controllers"
)

// @title UserService
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name Nekilc
// @contact.url http://www.nekilc.cn
// @contact.email nieaowei@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8020
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	r := gin.Default()
	hapi.RegisterHandleFunc(r, func(e *gin.Engine) *gin.RouterGroup {
		return e.Group("/api")
	},
		controllers.UserServiceController{},
		controllers.UnAuthController{},
		controllers.AuthController{},
	)
	cert,_ := tls.LoadX509KeyPair("./static/hfunc.nekilc.cn.pem","./static/hfunc.nekilc.cn.key")
	hapi.ServeAny(
		hhttp.WithAddr(":8020"),
		hhttp.WithHandler(r),
		hhttp.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h3", "h2", "h1"},
			ServerName:   "localhost",
		}),
	)
}
