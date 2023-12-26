package gateways

import (
	auth "notes-api-golang/adapter/gateways/auth"
	note "notes-api-golang/adapter/gateways/note"

	"github.com/gin-gonic/gin"
)

func CreateRoute(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	auth.AuthRouteGroup(v1)
	note.NoteRouteGroup(v1)

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
