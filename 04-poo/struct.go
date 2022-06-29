package main

import "time"

var Page struct {
	number int
	header string
	text   string
}

var Book struct {
	title        string
	author       string
	realeaseDate *time.Time
	pages        []struct {
		number int
		header string
		text   string
	}
}
