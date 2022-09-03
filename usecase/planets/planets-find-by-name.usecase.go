package planetsusecase

import (
	"fmt"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils"
	"github.com/BaianorASR/go-star-wars/utils/redis"
)

func (u *planetsUseCase) FindByName(query string) (*entities.Planet, error) {
	cacheKey := fmt.Sprintf("planet:%s", query)

	// Check if exists cached value
	if cached, err := redis.Get[entities.Planet](cacheKey); err != nil {
		return nil, err
	} else if cached != nil {
		return cached, nil
	}

	planet, err := u.repo.FindByName(query)
	if err != nil {
		return nil, err
	}

	length, err := utils.GetPlanetLength(planet.Name)
	if err != nil {
		return nil, err
	}
	planet.Films = length

	// Save in cache
	if err := redis.Set(cacheKey, planet); err != nil {
		return nil, err
	}

	return planet, nil
}
