package routes

import (
	planetscontroller "github.com/BaianorASR/go-star-wars/controller/planets"
	"github.com/BaianorASR/go-star-wars/middleware"
	"github.com/gin-gonic/gin"
)

func planetRoute(router *gin.RouterGroup) {
	r := router.Group("/planets")
	// Reset database
	r.DELETE("/reset-db", planetscontroller.Reset)

	// CRUD
	r.POST("/", middleware.PlanetsValidate, planetscontroller.Create)
	r.GET("/", planetscontroller.Find)
	r.GET("/search", planetscontroller.FindByName)
	r.GET("/:id", planetscontroller.FindById)
	r.DELETE("/:id", planetscontroller.DeleteOne)
}
