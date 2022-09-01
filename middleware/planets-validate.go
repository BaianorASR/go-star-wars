package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
)

type tovalidate struct {
	entities.Planet
}

var regex = regexp.MustCompile(`^[a-zA-Z]{0,255}$`)

func (t *tovalidate) Validate() error {
	// Name
	if reflect.ValueOf(t.Name).IsZero() {
		return errors.BadRequestf("The \"name\" field cannot be empty")
	}

	if !regex.MatchString(t.Name) {
		return errors.BadRequestf("The \"name\" field must be a string with a maximum of 255 characters")
	}

	if len(t.Name) < 3 {
		return errors.BadRequestf("The \"name\" field must be greater than 3")
	}

	// Climate
	if reflect.ValueOf(t.Climate).IsZero() {
		return errors.BadRequestf("The \"climate\" field cannot be empty")
	}

	if !regex.MatchString(t.Climate) {
		return errors.BadRequestf("The \"climate\" field must be a string with a maximum of 255 characters")
	}

	if len(t.Climate) < 3 {
		return errors.BadRequestf("The \"climate\" field must be greater than 3")
	}

	// Terrain
	if reflect.ValueOf(t.Terrain).IsZero() {
		return errors.BadRequestf("The \"terrain\" field cannot be empty")
	}

	if !regex.MatchString(t.Terrain) {
		return errors.BadRequestf("The \"terrain\" field must be a string with a maximum of 255 characters")
	}

	if len(t.Terrain) < 3 {
		return errors.BadRequestf("The \"terrain\" field must be greater than 3")
	}

	return nil
}

func PlanetsValidate(c *gin.Context) {
	ByteBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))

	var t tovalidate
	err := json.Unmarshal(ByteBody, &t)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("Invalid JSON: %s", err.Error()),
			})
		return
	}

	if err := t.Validate(); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
		return
	}

	c.Next()
}
