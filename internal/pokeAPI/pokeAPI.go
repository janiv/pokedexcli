package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func getLocationArea(offset string, limit string) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%s&limit=%s", offset, limit)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Got statuscode %d!", res.StatusCode)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", body)
}
