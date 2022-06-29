package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"01D9X58E": "Mad Jac Mixer",
	"01D9X5BQ": "Keystone Ice",
	"01D9X5CV": "Belgian Moon",
}

func main() {
	// b := flag.Bool("beers", false, "print beers")
	// flag.Parse()
	// if *b {
	// 	fmt.Println(beers)
	// }
	// fmt.Println(b)

	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specified a command beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find by ID")
		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}
}
