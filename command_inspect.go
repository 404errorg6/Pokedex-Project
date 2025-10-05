package main

import (
	"encoding/json"
	"fmt"
)

func commandInspect(_ *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("inspect command requires a pokemon name as argument")
	}

	body, ok := MyPokemons[pokemon].([]byte)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon\n")
	}

	var val map[string]any
	if err := json.Unmarshal(body, &val); err != nil {
		return fmt.Errorf("Error Unmarshalling data: %v", err)
	}

	name := val["name"].(string)
	height := val["height"].(float64)
	weight := val["weight"].(float64)
	stats := val["stats"].([]any)
	types := val["types"].([]any)

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Height: %v\n", height)
	fmt.Printf("Weight: %v\n", weight)

	fmt.Printf("Stats:\n")
	for _, stat := range stats {
		amount := stat.(map[string]any)["base_stat"].(float64)
		statVal := stat.(map[string]any)["stat"].(map[string]any)
		statName := statVal["name"].(string)
		fmt.Printf("	-%v: %v\n", statName, amount)
	}

	fmt.Printf("Types:\n")
	for _, pkType := range types {
		typeName := pkType.(map[string]any)["type"].(map[string]any)["name"].(string)
		fmt.Printf("	- %v\n", typeName)
	}

	return nil
}
