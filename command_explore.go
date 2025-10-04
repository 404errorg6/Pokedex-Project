package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(config *config, area string) error {
	if area == "" {
		return fmt.Errorf("explore requires a city as an argument")
	}

	fmt.Printf("Exploring %v...\n", area)
	url := URL + area

	var body []byte
	if val, ok := cache.Get(url); ok { // Cache hit
		body = val
	} else { // Cache miss
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("Failed with status code %d : %v\n", res.StatusCode, string(body))
		}
		if err != nil {
			return err
		}
		cache.Add(url, body)
	}

	pokemonData := map[string]any{}

	if err := json.Unmarshal(body, &pokemonData); err != nil {
		return err
	}

	pokemonEncounters, ok := pokemonData["pokemon_encounters"].([]any)
	if !ok {
		return fmt.Errorf("Unexpected response: Could not find pokemon_encounters field")
	}

	fmt.Printf("Found Pokemon: \n")
	for _, pkEncounter := range pokemonEncounters {

		pokemon, ok := pkEncounter.(map[string]any)["pokemon"].(map[string]any)
		if !ok {
			return fmt.Errorf("Unexpected response: Could not find pokemon field\n")
		}

		name, ok := pokemon["name"].(string)
		if !ok {
			return fmt.Errorf("Unexpected response: Could not find name field\n")
		}
		fmt.Printf("	- %v\n", name)
	}
	return nil
}
