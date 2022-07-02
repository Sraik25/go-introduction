package inmen

import beerscli "github.com/Sraik25/golang-introduction/08-automated_tests/internal"

type repository struct {
}

func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers implements beerscli.BeerRepo
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	return []beerscli.Beer{
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Domestic Specialty",
			"Molson",
			"Canada",
			beerscli.NewBeerType("Lager"),
			23.95,
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Grolsch Export B.V.",
			"Canada",
			beerscli.NewBeerType("Non-Alcoholic Beer"),
			49.50,
		),
	}, nil
}
