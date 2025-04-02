package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "CHARMANDER BULBASAUR SQUIRTLE",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "CHARMANDER BULBASAUR   squirtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "       ",
			expected: []string{},
		},
		{
			input:    "yer a wizard,                harry",
			expected: []string{"yer", "a", "wizard,", "harry"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %s; got: %s", expectedWord, word)
			}
		}
	}

}
