package main

import (
	"auth-service/api/http/controllers"
	"crypto/tls"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
)

//go:embed static
var static embed.FS

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"msg": "OK",
		})
	})
	hapi.RegisterHandleFunc(
		r,
		func(e *gin.Engine) *gin.RouterGroup {
			return e.Group("/api")
		},
		controllers.AuthServiceController{},
	)
	cert,_ := tls.LoadX509KeyPair("./static/hfunc.nekilc.cn.pem","./static/hfunc.nekilc.cn.key")
	hapi.ServeAny(
		hhttp.WithAddr(":8010"),
		hhttp.WithHandler(r),
		hhttp.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h3", "h2", "h1"},
			ServerName:   "localhost",
		}),
	)
}
