package planetscontroller

import (
	"net/http"
	"strings"

	"github.com/BaianorASR/go-star-wars/database"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
)

func FindByName(c *gin.Context) {
	// Get query
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query is required",
			"status":  http.StatusBadRequest,
		})
		return
	}

	q := strings.ToUpper(query[:1]) + strings.ToLower(query[1:])

	// Integration with usecase
	db := database.ConnectDB()
	defer db.Disconnect(c)

	repo := planetsrepository.New(db)
	useCase := planetsusecase.New(repo)

	// Find planet
	planet, err := useCase.FindByName(q)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		}
		return
	}

	// Return planet
	c.JSON(http.StatusOK, planet)
}
