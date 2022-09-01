package planetscontroller

import (
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteOne controller calls the usecase to delete a planet
func DeleteOne(c *gin.Context) {
	var err error

	// make implementation
	Db := database.ConnectDB()
	repo := planetsrepository.New(Db)
	useCase := planetsusecase.New(repo)

	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	err = useCase.DeleteOne(objId)
	if err != nil {
		if errors.IsNotFound(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":  err.Error(),
				"status": http.StatusNotFound,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "Planet deleted",
		"status":  http.StatusOK,
	})

}
