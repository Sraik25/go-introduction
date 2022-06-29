package main

import "time"

type page struct {
	number int
	header string
	text   string
}

type book struct {
	title        string
	author       string
	realeaseDate *time.Time
	pages        []page
}
