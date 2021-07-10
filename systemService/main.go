package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	controller2 "system-service/api/http/controller"
)

func main() {
	hapi.Server("8000", nil, func(c *gin.Engine) {
		hapi.RegisterHandleFunc(c, func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/api")
		}, controller2.SystemServiceController{})
	})
}
