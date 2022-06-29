package main

import "fmt"

func main() {
	b := make(map[string]float32)

	b["mahou"] = 0.59
	b["buckler"] = 0.60

	fmt.Println("beers: ", b)

	b1 := b["mahou"]
	fmt.Println("beer 1: ", b1)

	delete(b, "buckler")
	fmt.Println("beers: ", b)
}
