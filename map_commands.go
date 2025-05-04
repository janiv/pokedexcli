package main

import (
	"fmt"
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
	fmt.Println(args[1])
	return nil
}
