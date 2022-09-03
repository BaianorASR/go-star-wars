package planetsrepository

import (
	"context"
	"reflect"
	"strings"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/americanas-go/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *planetsRepository) Create(planet *entities.Planet) (*entities.Planet, error) {
	var exists entities.Planet
	r.client.FindOne(context.Background(), bson.M{"name": planet.Name}).Decode(&exists)

	if !reflect.ValueOf(exists).IsZero() {
		return nil, errors.NewConflict(nil, "planet already exists")
	}

	result, err := r.client.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: planet.Name},
		{Key: "climate", Value: planet.Climate},
		{Key: "terrain", Value: planet.Terrain},
		{Key: "films", Value: planet.Films},
	})
	if err != nil {
		if strings.Contains(err.Error(), "Unauthorized") {
			return nil, errors.Unauthorizedf("Unauthorized to create planet, please see you credentials")
		}
		return nil, errors.Internalf(err.Error())
	}

	planet.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return planet, nil
}
