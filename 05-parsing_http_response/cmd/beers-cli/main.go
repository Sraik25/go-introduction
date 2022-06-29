package main

import (
	"github.com/Sraik25/golang-introduction/05-parsing_http_response/internal/cli"
	"github.com/Sraik25/golang-introduction/05-parsing_http_response/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.Execute()
}
