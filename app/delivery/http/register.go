package http

import (
	"github.com/DarkSoul94/golang-template/app"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc app.Usecase) {
	h := NewHandler(uc)

	apiEndpoints := router.Group("/api")
	{
		apiEndpoints.POST("/", h.HelloWorld)
	}
}
