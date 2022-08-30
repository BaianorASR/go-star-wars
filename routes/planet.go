package routes

import "github.com/gin-gonic/gin"

func planetRoute(router *gin.RouterGroup) {
	r := router.Group("/talker")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
