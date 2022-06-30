package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/Sraik25/golang-introduction/06-error_handling/internal"
	"github.com/Sraik25/golang-introduction/06-error_handling/internal/errors"
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
		return nil, &errors.BadResponseErr{Msg: "Something happened when call the endpoint", File: "ontario/repository.go", Line: 31}
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &errors.BadResponseErr{Msg: "Something happened when read the content", File: "ontario/repository.go", Line: 36}
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, &errors.BadResponseErr{Msg: "Something happened when unmarshal the content", File: "ontario/repository.go", Line: 41}
	}
	return
}
