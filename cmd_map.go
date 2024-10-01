package main

import(
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"
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

	cachedData, ok := config.Cache.Get(url) 
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

		config.Cache.Add(url, body)
	
		var location LocationMap
	
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
			fmt.Println(r.Name)
		}
	} else {
		//TODO: IMPLEMENT CACHING

		var location LocationMap

		err := json.Unmarshal(cachedData, &location)
        if err != nil {
            return err
        }

		for i, r := range location.Results {
            if i >= 20 {
                break
            }
            fmt.Println(r.Name)
        }
	}
	return nil
}

type LocationMap struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}