package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/Sraik25/golang-introduction/07-behaviour_error_handling/internal"
	"github.com/Sraik25/golang-introduction/07-behaviour_error_handling/internal/errors"
)

const (
	productsEnpoint = "/products"
	ontarioURL      = "http://localhost:4000"
)

type beerRepo struct {
	url string
}

// new
func NewOntarioRepository() beerscli.BeerRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEnpoint))

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", productsEnpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading to response from %s", productsEnpoint)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}
