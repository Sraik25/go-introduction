package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/Sraik25/golang-introduction/05-parsing_http_response/internal"
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
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}
	return
}
