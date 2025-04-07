package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := NewConfig()
	coms := getCommands()
	for {
		fmt.Print("Pokedex ->")
		scanner.Scan()
		user_input := scanner.Text()
		user_slice := cleanInput(user_input)
		given_command := user_slice[0]
		val, ok := coms[given_command]
		if ok {
			val.callback(&cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
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
