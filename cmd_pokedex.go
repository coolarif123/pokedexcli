package main

import "fmt"

func commandPokedex(config *Config) error {
	if len(config.PokemonCaught) == 0 {
		fmt.Println("Damn you suck. Catch something ffs.")
	}

	for _, r := range config.PokemonCaught {
		fmt.Printf("- %s\n", r.Name)
	}
	return nil
}