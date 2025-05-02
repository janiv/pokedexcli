package pokeapi

import (
	"fmt"
	"testing"
)

func TestMapAPI(t *testing.T) {
	cases := []struct {
		key string
		val string
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			es := NewExtraStruct()
			_, next, _, err := es.MapAPI(c.key)
			if err != nil {
				t.Errorf("Map api returned error")
				return
			}
			if string(next) != string(c.val) {
				t.Errorf("wrong .next value, got %s instead of %s", next, c.val)
				return
			}

		})
	}
}
