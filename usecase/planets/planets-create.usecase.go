package planetsusecase

import (
	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils/redis"
	"github.com/americanas-go/errors"
)

// Create a new planet
func (u *planetsUseCase) Create(planet *entities.Planet) (*entities.Planet, error) {
	planet, err := u.repo.Create(planet)
	if err != nil {
		return nil, errors.NewInternal(err, "error on create")
	}

	// Get all cached planets
	planets, err := redis.Get[[]entities.Planet]("planets-find-all")
	if err != nil {
		return nil, err
	}

	// Add new planet to cache
	if planets != nil {
		*planets = append(*planets, *planet)
		if err := redis.Set("planets-find-all", planets); err != nil {
			return nil, err
		}
	}

	return planet, nil
}
