package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Res struct {
	Next   string   `json:"next" bson:"next"`
	Result []planet `json:"results" bson:"results"`
}

func getPlanets() []planet {
	var planets Res

	response, err := http.Get("https://swapi.dev/api/planets/")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(responseData, &planets)

	if planets.Next != "" {
		println("NEXT")
		for planets.Next != "" {
			println(planets.Next)

			response, err = http.Get(planets.Next)
			if err != nil {
				fmt.Println(err)
			}

			responseData, err = ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			var planetsNext Res

			json.Unmarshal(responseData, &planetsNext)

			planets.Result = append(planets.Result, planetsNext.Result...)
			planets.Next = planetsNext.Next
			println("END")
		}
	}

	return planets.Result
}

// insetMany inserts the planets in the database
func insetMany() {
	fmt.Println("START")

	planets := getPlanets()

	fmt.Println(len(planets))
	fmt.Println(planets)

	db := database.ConnectDB().Database("go-americanas-teste").Collection("planets")
	for _, p := range planets {
		_, err := db.InsertOne(context.Background(), bson.D{
			{Key: "name", Value: p.Name},
			{Key: "climate", Value: p.Climate},
			{Key: "terrain", Value: p.Terrain},
		})
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("END INSERT")
}

// deletePlanets deletes all planets in the database
func deletePlanets() {
	db := database.ConnectDB().Database("go-americanas-teste").Collection("planets")

	_, err := db.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
}

func Reset() {
	deletePlanets()
	insetMany()
}
