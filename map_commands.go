package main

import (
	"fmt"
	"strings"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config, args ...string) error {
	res, next, prev, err := cfg.ES.MapAPI(cfg.Next)
	if err != nil {
		fmt.Println("Something went wrong!")
		return err
	}
	for i := range res {
		fmt.Println(res[i])
	}
	cfg.Next = next
	cfg.Previous = prev
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == "" {
		fmt.Println("You are on the first page")
		return nil
	}
	res, next, prev, err := cfg.ES.MapAPI(cfg.Previous)
	if err != nil {
		fmt.Println("Something went wrong!")
		return err
	}
	for i := range res {
		fmt.Println(res[i])
	}
	cfg.Next = next
	cfg.Previous = prev
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	base_url := strings.Split(cfg.Next, "?")[0]
	url := base_url + "/" + args[1]
	res, err := cfg.ES.Explore(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Found Pokemon:")
	for i := range res {
		fmt.Printf("  - %s\n", res[i])
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	base_url := "https://pokeapi.co/api/v2/pokemon/"
	pokemon := args[1]
	url := base_url + pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	res, err := cfg.ES.Catch(url, pokemon)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res {
		fmt.Printf("%s was caught!\n", pokemon)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
}
