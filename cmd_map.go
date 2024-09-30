package main

import(
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"
	"github.com/coolarif123/pokedexcli/internal/pokecache"
)

func commandMap(config *Config) error {
	var url string

	if config.initMap == false {
		url = "https://pokeapi.co/api/v2/location-area/"
		config.initMap = true
	}	else if config.Next != nil && !config.Mapb {
		url = *config.Next
	}	else if config.Previous != nil && config.Mapb {
		url = *config.Previous
	}	else {
		return errors.New("Error in retrieving next URL")
	}

	cachedData, ok := cache.data[url] 
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}		
		if res.StatusCode != 200 {
			return errors.New("Pokemon API not available")
		}
	
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		
	
		if err != nil {
			return err
		}
	
		var location Location
	
		err = json.Unmarshal(body, &location)
		if err != nil {
			return err
		}

		if location.Next != nil {
			config.Next = location.Next
		}	else {
			config.Next = nil
		}
		
		if location.Previous != nil {
			config.Previous = location.Previous
		} else {
			config.Previous = nil
		}

		for i, r := range location.Results {
			if i >= 20 {
				break
			} 
			cache.Add(url, []byte(r))
			fmt.Println(r.Name)
		}
	} else {
		//TODO: IMPLEMENT CACHING
		for i, r := range cachedData {
			if i >= 20 {
				break
			}
			fmt.Println(r)
		}
	}
	return nil
}

type Config struct {
	Mapb      bool
	initMap   bool
	Previous *string
	Next 	 *string
}

type Location struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}