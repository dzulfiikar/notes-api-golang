package gateways

import (
	route "notes-api-golang/adapter/gateways/routes"

	"github.com/gin-gonic/gin"
)

func CreateRoute(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	route.AuthRouteGroup(v1)
	route.NoteRouteGroup(v1)

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
