package planetsusecase

import (
	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID returns a planet by ID from the repository and the length of Star Wars API
func (u planetsUseCase) FindById(objId primitive.ObjectID) (*entities.Planet, error) {
	planet, err := u.repo.FindById(objId)
	if err != nil {
		return nil, err
	}

	if length, err := utils.GetPlanetLength(planet.Name); err != nil {
		return nil, err
	} else {
		planet.Films = length
	}

	return planet, nil
}
