package cli

import (
	"fmt"
	"log"
	"strconv"

	beerscli "github.com/Sraik25/golang-introduction/06-error_handling/internal"
	"github.com/Sraik25/golang-introduction/06-error_handling/internal/errors"
	"github.com/spf13/cobra"
)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(repository beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}
	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beers")

	return beersCmd
}

func runBeersFn(repository beerscli.BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		beers, err := repository.GetBeers()
		if err, ok := err.(*errors.BadResponseErr); ok {
			log.Fatal(err)
		}

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ProductID == i {
					fmt.Println(beer)
					return
				}
			}
		} else {
			fmt.Println(beers)
		}
	}
}
