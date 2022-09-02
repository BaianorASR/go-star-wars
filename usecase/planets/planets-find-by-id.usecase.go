package planetsusecase

import (
	"fmt"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils"
	"github.com/BaianorASR/go-star-wars/utils/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID returns a planet by ID from the repository and the length of Star Wars API
func (u planetsUseCase) FindById(objId primitive.ObjectID) (*entities.Planet, error) {
	cacheKey := fmt.Sprintf("planet:%s", objId)

	// Check if exists cached value
	if cached, err := redis.Get[entities.Planet](cacheKey); err != nil {
		return nil, err
	} else if cached != nil {
		return cached, nil
	}

	planet, err := u.repo.FindById(objId)
	if err != nil {
		return nil, err
	}

	if length, err := utils.GetPlanetLength(planet.Name); err != nil {
		return nil, err
	} else {
		planet.Films = length
	}

	if err := redis.Set(cacheKey, planet); err != nil {
		return nil, err
	}

	return planet, nil
}
