package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var res []string
	// First split on empty space
	slc := strings.Split(text, " ")
	for i := range slc {
		// Remove empty space
		slc[i] = strings.TrimSpace(slc[i])
		// Set to lower case
		slc[i] = strings.ToLower(slc[i])
		// Res should only have strings that are not empty
		if len(slc[i]) > 0 {
			res = append(res, slc[i])
		}
	}
	return res
}
