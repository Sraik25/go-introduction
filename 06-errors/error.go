package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	openFile()
}

func openFile() {
	_, err := os.Open("errors.mp4")
	if err != nil {
		panic(err)
	}
}
