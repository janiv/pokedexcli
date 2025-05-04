package main

import (
	"fmt"
	"strings"
)

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

func commandInspect(cfg *config, args ...string) error {
	pokemonName := args[1]
	in_dex, res := cfg.ES.Inspect(pokemonName)
	if !in_dex {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	} else {
		fmt.Printf("Name: %s\n", res.Name)
		fmt.Printf("Height: %d\n", res.Height)
		fmt.Printf("Weight: %d\n", res.Weight)
		fmt.Printf("Stats:\n")
		fmt.Printf("  -hp:%d\n", res.Hp)
		fmt.Printf("  -attack:%d\n", res.Attack)
		fmt.Printf("  -defense:%d\n", res.Defense)
		fmt.Printf("  -special-attack:%d\n", res.SpecialAttack)
		fmt.Printf("  -special-defense:%d\n", res.SpecialDefense)
		fmt.Printf("  -speed:%d\n", res.Speed)
		fmt.Printf("Types:\n")
		fmt.Printf("  -%s\n", res.Type1)
		if res.Type2 != "" {
			fmt.Printf("  -%s\n", res.Type2)
		}
		return nil
	}
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.ES.Pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
