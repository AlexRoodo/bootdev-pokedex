package main

import (
	"fmt"
	"strings"
)

func commandHelp(cfg *config, args ...string) error {
	var result strings.Builder
	result.WriteString("\nWelcome to the Pokedex!\nUsage:\n")
	for _, c := range getCommands() {
		result.WriteString("\n")
		result.WriteString(c.name)
		result.WriteString(": ")
		result.WriteString(c.description)
	}
	fmt.Println(result.String())
	return nil
}
