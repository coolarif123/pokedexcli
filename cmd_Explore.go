package main

import (
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"
	// "strings"
)
//TODO:
// - Check if area input exists or not if 200 not returned then the area input does not exist
// - 

func commandExplore(config *Config) error {
	const urlBase = "https://pokeapi.co/api/v2/location-area/"
	url := urlBase + config.Area

	cachedData, ok := config.Cache.Get(url) 
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}		
		if res.StatusCode != 200 {
			return errors.New("Location does not exist")
		}
	
		fmt.Println("Exploring " + config.Area + "...")

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
	
		if err != nil {
			return err
		}

		config.Cache.Add(url, body)
	
		var location Location

		err = json.Unmarshal(body, &location)
		if err != nil {
			return err
		}

		if len(location.PokemonEncounters) == 0  {
			fmt.Println("No Pokemon Found in this area")
		} else {
			fmt.Println("Found Pokemon: ")
			for _, encounter := range location.PokemonEncounters {
				fmt.Println(encounter.Pokemon.Name)
			}
		}
	} else {
		//TODO: IMPLEMENT CACHING

		var location Location

		err := json.Unmarshal(cachedData, &location)
        if err != nil {
            return err
        }

		if len(location.PokemonEncounters) == 0  {
			fmt.Println("No Pokemon Found in this area")
		} else {
			fmt.Println("Found Pokemon: ")
			for _, encounter := range location.PokemonEncounters {
				fmt.Println(encounter.Pokemon.Name)
			}
		}
	}
	return nil
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}