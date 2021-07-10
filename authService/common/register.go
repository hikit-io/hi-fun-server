package common

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
}

type HandleFunc func() (httpMethod, routeUri, version string, handlerFunc fiber.Handler)
type NewHandleFunc func() HandleFunc

var _EmptyHandleFunc = HandleFunc(
	func() (httpMethod, routeUri, version string, handleFunc fiber.Handler) {
		return "", "", "", nil
	},
)
var _EmptyNewHandleFunc = NewHandleFunc(
	func() HandleFunc {
		return func() (httpMethod, routeUri, version string, handleFunc fiber.Handler) {
			return "", "", "", nil
		}
	},
)

func RegisterController(r fiber.Router, c Controller) {

}
