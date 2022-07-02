package main

import (
	"flag"

	beerscli "github.com/Sraik25/golang-introduction/08-automated_tests/internal"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/cli"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/fetching"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/storage/csv"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/storage/ontario"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo

	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.Execute()
}
