package planetsusecase

import (
	"sync"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/BaianorASR/go-star-wars/utils"
)

func (u planetsUseCase) Find(query string) (*[]entities.Planet, error) {
	planets, err := u.repo.Find(query)
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

	return planets, nil
}
