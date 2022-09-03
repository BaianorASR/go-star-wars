package planetsusecase

import (
	"fmt"
	"sync"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils"
	"github.com/BaianorASR/go-star-wars/utils/redis"
)

// Find return all planets
func (u planetsUseCase) Find() (any, error) {
	cacheKey := "planets-find-all"

	// Check if exists cached value
	if cached, err := redis.Get[[]entities.Planet](fmt.Sprintf(cacheKey)); err != nil {
		return nil, err
	} else if cached != nil {
		return cached, nil
	}

	planets, err := u.repo.Find()
	if err != nil {
		return nil, err
	}

	// Create a wait group
	wg := sync.WaitGroup{}
	wg.Add(len(*planets))

	for i, p := range *planets {
		// Create a goroutine for each planet
		go func(i int, p entities.Planet) {
			defer wg.Done()

			length, err := utils.GetPlanetLength(p.Name)
			if err != nil {
				return
			}

			(*planets)[i].Films = length
		}(i, p)

		// Wait for all goroutines to finish
		if i == len(*planets)-1 {
			wg.Wait()
		}
	}

	if err := redis.Set(cacheKey, planets); err != nil {
		return nil, err
	}

	return planets, nil
}
