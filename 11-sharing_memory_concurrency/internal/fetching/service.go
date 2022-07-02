package fetching

import (
	"math"
	"sync"

	beerscli "github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal"
)

type Service interface {
	FetchBeers() ([]beerscli.Beer, error)
	FetchByID(id int) (beerscli.Beer, error)
}

type service struct {
	bR beerscli.BeerRepo
}

func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

// FetchBeers implements Service
func (s *service) FetchBeers() ([]beerscli.Beer, error) {
	return s.bR.GetBeers()
}

// FetchByID implements Service
func (s *service) FetchByID(id int) (beerscli.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return beerscli.Beer{}, err
	}

	beersPerRoutine := 10
	numRoutines := numOfRoutines(len(beers), beersPerRoutine)

	wg := &sync.WaitGroup{}

	wg.Add(numRoutines)

	var b beerscli.Beer

	for i := 0; i < numRoutines; i++ {
		go func(id, begin, end int, beers []beerscli.Beer, b *beerscli.Beer, wg *sync.WaitGroup) {
			for i := begin; i < end; i++ {
				if beers[i].ProductID == id {
					*b = beers[i]
				}
			}
			wg.Done()
		}(id, i, i+beersPerRoutine, beers, &b, wg)
	}

	wg.Wait()

	return b, nil
}

func numOfRoutines(numOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numOfBeers) / float64(beersPerRoutine)))
}
