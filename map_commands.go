package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func commandMap(cfg *config) error {
	resp, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	la := LocationAreas{}
	json_err := json.Unmarshal(body, &la)
	if json_err != nil {
		return json_err
	}
	for i := range la.Results {
		fmt.Println(la.Results[i].Name)
	}
	cfg.Previous = la.Previous
	cfg.Next = la.Next

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	la := LocationAreas{}
	json_err := json.Unmarshal(body, &la)
	if json_err != nil {
		return json_err
	}
	for i := range la.Results {
		fmt.Println(la.Results[i].Name)
	}
	cfg.Next = la.Next
	cfg.Previous = la.Previous

	return nil
}
