package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(config *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("catch command requires a pokemon name as argument.\n")
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	url := URL + "pokemon/" + pokemon
	var body []byte

	if val, ok := cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		val, err := io.ReadAll(res.Body)
		body = val
		res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("Error occured with status code %v: %v\n", res.StatusCode, string(body))
		}
		if err != nil {
			return err
		}
		cache.Add(url, body)
	}

	pkDetails := map[string]any{}

	if err := json.Unmarshal(body, &pkDetails); err != nil {
		return err
	}

	baseExp := int(pkDetails["base_experience"].(float64))
	chance := rand.Intn(baseExp)

	if chance == 0 {
		chance = 1
	}

	if true {
		MyPokemons[pokemon] = body
		fmt.Printf("%v was caught\n", pokemon)
		return nil
	}

	fmt.Printf("%v escaped!\n", pokemon)
	return nil
}
