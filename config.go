package main

import (
	"github.com/janiv/pokedexcli/internal/pokeapi"
)

type config struct {
	Next     string
	Previous string
	ES       pokeapi.ExtraStruct
}

func newConfig() *config {
	return &config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
		ES:       *pokeapi.NewExtraStruct(),
	}
}
