package main

import (
	"fmt"
)
//TODO:
// - Check if area input exists or not if 200 not returned then the area input does not exist
// - 

func commandExplore(area string) error {
	const urlBase := "https://pokeapi.co/api/v2/location-area/"


}

type PokemonInArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	}
}