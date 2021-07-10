package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
)

type SystemServiceController struct {
}

func (s SystemServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s SystemServiceController) RouterGroupName() (name string) {
	return "system_service"
}

func (s SystemServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s SystemServiceController) Version() string {
	return "v1"
}

func (s SystemServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}

type CarouselResp struct {
	Count int64 `json:"count"`
	List  []struct {
		Id  int64  `json:"id"`
		Url string `json:"url"`
		Img string `json:"img"`
	} `json:"list"`
}

// https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999
func (s SystemServiceController) CarouselList() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "carousel/list", "v1", func(ctx *gin.Context) {
		var (
			resp = new(CarouselResp)
		)
		resp.Count = 3
		resp.List = []struct {
			Id  int64  `json:"id"`
			Url string `json:"url"`
			Img string `json:"img"`
		}{
			{
				1,
				"",
				"https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999",
			},
			{
				2,
				"",
				"https://ali-cdn.educoder.net/images/avatars/PortalImage/63?t=1613987218",
			},
			{
				3,
				"",
				"https://ali-cdn.educoder.net/images/avatars/PortalImage/96?t=1616488300",
			},
		}
		hapi.JsonResponseOk(ctx, resp)
	}
}

type BannerResp struct {
	Url string `json:"url"`
}

type BannerReq struct {
	Path string `form:"path" binding:"required"`
}

// https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999
func (s SystemServiceController) BannerOne() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "banner/one", "v1", func(context *gin.Context) {
		var (
			err    error
			params = new(BannerReq)
			resp   = new(BannerResp)
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(context, err)
			}
		}()
		if err = context.ShouldBind(params); err != nil {
			return
		}
		switch params.Path {
		case "/index":
			err = hapi.NewCodeErr(2, "")
			return
		case "/practice":
			resp.Url = "https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999"
		case "/competition":
			resp.Url = "https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999"
		case "/about":
			resp.Url = "https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999"
		}
		hapi.JsonResponseOk(context, resp)
	}
}

type AnnounceItem struct {
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
	Labels  []string `json:"labels"`
	Content string   `json:"content"`
}

type AnnounceListResp struct {
	Count int64          `json:"count"`
	List  []AnnounceItem `json:"list"`
}

type AnnounceListReq struct {
	Limit int64 `json:"limit"`
}

// https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999
func (s SystemServiceController) AnnounceList() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "announce/list", "v1", func(context *gin.Context) {
		var (
			err    error
			params = new(AnnounceListReq)
			resp   = new(AnnounceListResp)
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(context, err)
			} else {
				hapi.JsonResponseOk(context, resp)
			}
		}()
		if err = context.ShouldBind(params); err != nil {
			return
		}
		resp.List = []AnnounceItem{
			{
				Id:      1,
				Title:   "十点吃饭",
				Content: "# 十点吃饭",
			},
			{
				Id:      2,
				Title:   "九点吃饭",
				Content: "# 九点吃饭",
			},
			{
				Id:      3,
				Title:   "八点吃饭",
				Content: "# 八点吃饭",
			},
		}
		resp.Count = 3
		return
	}
}

type AnnounceOneResp struct {
	AnnounceItem
}

type AnnounceOneReq struct {
	Id int64 `json:"id"`
}

func (s SystemServiceController) AnnounceOne() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "announce/one", "v1", func(context *gin.Context) {
		var (
			err    error
			params = new(AnnounceOneReq)
			resp   = new(AnnounceOneResp)
		)
		defer func() {
			if err != nil {
				hapi.JsonResponseErr(context, err)
			} else {
				hapi.JsonResponseOk(context, resp)
			}
		}()
		if err = context.ShouldBind(params); err != nil {
			return
		}
		resp.Id = 1
		resp.Content = "## 吃饭"
		return
	}
}
