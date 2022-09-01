package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/americanas-go/errors"
)

type planet struct {
	ID      string   `json:"id" bson:"_id"`
	Name    string   `json:"name" bson:"name"`
	Climate string   `json:"climate" bson:"climate"`
	Terrain string   `json:"terrain" bson:"terrain"`
	Films   []string `json:"films" bson:"films"`
}

type res struct {
	Next   string   `json:"next" bson:"next"`
	Result []planet `json:"results" bson:"results"`
}

// GetPlanetLength returns the length of the planets in the Star Wars API
func GetPlanetLength(name string) (int, error) {
	var err error

	response, err := http.Get(fmt.Sprintf("https://swapi.dev/api/planets/?search=%s&format=json", name))
	if err != nil {
		return 0, errors.Internalf("Error on get the length of the planet: %s", err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, errors.Internalf("Error on parse body of the response: %s", err.Error())
	}

	var p res
	json.Unmarshal(responseData, &p)

	if len(p.Result) == 0 {
		return 0, nil
	}

	return len(p.Result[0].Films), nil
}
