package planetscontroller

import (
	"github.com/BaianorASR/go-star-wars/utils"
	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {
	utils.Reset()
}
