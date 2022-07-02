package fetching

import (
	"math"

	beerscli "github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal"
	"github.com/pkg/errors"
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

	b := make(chan beerscli.Beer)
	done := make(chan bool, numRoutines)

	for i := 0; i < numRoutines; i++ {
		toSearch := make([]beerscli.Beer, beersPerRoutine)
		copy(beers[i:i+beersPerRoutine], toSearch[:])

		go func(beers []beerscli.Beer, b chan beerscli.Beer, done chan bool) {
			for _, beer := range beers {
				if beer.ProductID == id {
					b <- beer
				}
			}
			done <- true
		}(toSearch, b, done)

	}

	var beer beerscli.Beer
	i := 0
	for i < numRoutines {
		select {
		case beer := <-b:
			return beer, nil

		case <-done:
			i++
		}

	}

	return beer, errors.Errorf("Beer %d not found", id)
}

func numOfRoutines(numOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numOfBeers) / float64(beersPerRoutine)))
}
