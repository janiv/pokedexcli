package main

type config struct {
	Next     string
	Previous string
}

func newConfig() *config {
	return &config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
}
