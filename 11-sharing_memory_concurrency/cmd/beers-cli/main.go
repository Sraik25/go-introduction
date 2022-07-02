package main

import (
	"flag"

	beerscli "github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal"
	"github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal/cli"
	"github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal/fetching"
	"github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal/storage/csv"
	"github.com/Sraik25/golang-introduction/11-sharing_memory_concurrency/internal/storage/ontario"
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
