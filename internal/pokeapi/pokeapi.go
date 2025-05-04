package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/janiv/pokedexcli/internal/pokecache"
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

type LocationAreaByName struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type ExtraStruct struct {
	client http.Client
	cache  pokecache.Cache
}

func NewExtraStruct() *ExtraStruct {
	var es ExtraStruct
	es.client = http.Client{}
	dur, _ := time.ParseDuration("10s")
	es.cache = *pokecache.NewCache(dur)
	return &es
}

func (es *ExtraStruct) MapAPI(url string) ([]string, string, string, error) {
	val, ok := es.cache.Get(url)
	la := LocationAreas{}
	if ok {
		json_err := json.Unmarshal(val, &la)
		if json_err != nil {
			return nil, "", "", json_err
		}
		res := make([]string, 20)
		for i := range la.Results {
			res[i] = la.Results[i].Name
		}
		fmt.Printf("%s was in cache!\n", url)
		return res, la.Next, la.Previous, nil
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return nil, "", "", err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, "", "", err
		}
		es.cache.Add(url, body)
		json_err := json.Unmarshal(body, &la)
		if json_err != nil {
			return nil, "", "", json_err
		}
		res := make([]string, 20)
		for i := range la.Results {
			res[i] = la.Results[i].Name
		}
		return res, la.Next, la.Previous, nil
	}
}

func (es *ExtraStruct) Explore(url string) ([]string, error) {
	val, ok := es.cache.Get(url)
	lan := LocationAreaByName{}
	if ok {
		json_err := json.Unmarshal(val, &lan)
		if json_err != nil {
			return nil, json_err
		}
		res := make([]string, 0)
		for i := range lan.PokemonEncounters {
			res = append(res, lan.PokemonEncounters[i].Pokemon.Name)
		}
		fmt.Printf("%s was in cache!\n", url)
		return res, nil
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		es.cache.Add(url, body)
		json_err := json.Unmarshal(body, &lan)
		if json_err != nil {
			return nil, json_err
		}
		res := make([]string, 0)
		for i := range lan.PokemonEncounters {
			res = append(res, lan.PokemonEncounters[i].Pokemon.Name)
		}
		return res, nil
	}
}
