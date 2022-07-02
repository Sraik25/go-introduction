package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal"
	"github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal/errors"
	jsoniter "github.com/json-iterator/go"
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

	err = b.betterdUnmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}

func (b *beerRepo) standardUnmarshal(data []byte, beers *[]beerscli.Beer) error {
	err := json.Unmarshal(data, &beers)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return nil
}

func (b *beerRepo) betterdUnmarshal(data []byte, beers *[]beerscli.Beer) error {
	var js = jsoniter.ConfigCompatibleWithStandardLibrary

	err := js.Unmarshal(data, &beers)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}

	return nil
}
