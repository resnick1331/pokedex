package main

import (
	"strings"
)

// Functions
// 01
func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowerCased := strings.ToLower(trimmed)
	words := strings.Fields(lowerCased)

	return words
}