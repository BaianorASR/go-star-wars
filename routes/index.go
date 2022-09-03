package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	planetRoute(superRoute)
	// loginRoute(superRoute)
}
 