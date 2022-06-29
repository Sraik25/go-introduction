package main

import "fmt"

func main() {
	g := gopher{animal{"Olek's Gopher"}, 4}
	g.Print()
}

type animal struct {
	name string
}

func (a animal) Print() {
	fmt.Printf("I'm %s and I'm a %T\n", a.name, a)
}

type gopher struct {
	animal
	legs int
}

func (g gopher) Print() {
	fmt.Printf("I'm %s and I'm a %T\n", g.name, g)
}
