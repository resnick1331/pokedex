package main

import (

	"testing"
)

func TestCleanInput (t *testing.T){
		cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  multiple   spaces  ",
			expected: []string{"multiple", "spaces"},
		},
		{
			input:    "no extra spaces",
			expected: []string{"no", "extra", "spaces"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: %v vs %v", actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) word at index %d is %q, want %q", c.input, i, actual[i], c.expected[i])
			}

		}
	}

}