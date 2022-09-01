package planetscontroller

import (
	"fmt"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	"github.com/BaianorASR/go-star-wars/entities"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := database.ConnectDB()
	repo := planetsrepository.New(db)
	useCase := planetsusecase.New(repo)

	var planet *entities.Planet
	c.BindJSON(&planet)

	fmt.Println(planet)

	planet, err := useCase.Create(planet)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"planet": planet,
	})
}
