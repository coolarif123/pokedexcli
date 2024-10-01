package main

import "fmt"

func commandInspect(config *Config) error {
	if pokemon, exists := config.PokemonCaught[config.PokemonName]; !exists {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, r := range pokemon.Stats{
			fmt.Printf("-%s: %d\n", r.Stat.Name, r.BaseStat)
		}
		fmt.Println("Types: ")
		for _, r := range pokemon.Types {
			fmt.Printf("- %s\n", r.Type.Name)
		}
	}
	return nil
}