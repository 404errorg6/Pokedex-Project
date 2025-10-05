package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/404errorg6/Pokedex-Project/internal"
)

var (
	MyPokemons = map[string]any{}
	URL        = "https://pokeapi.co/api/v2/"
	cfg        config
	cache      = internal.NewCache(time.Second * 5)
)

type config struct {
	Count   int     `json:"count"`
	NextURL string  `json:"next"`
	PrevURL *string `json:"previous"`
	Results []names `json:"results"`
}

type names struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func implementCache(key string) bool {
	var check bool

	if data, ok := cache.Map[key]; ok {

		if err := json.Unmarshal(data.Val, &cfg); err != nil {
			fmt.Printf("Error Unmarhsalling in cache: %v", err)
			return check
		}

		for _, result := range cfg.Results {
			fmt.Printf("%v\n", result.Name)
			check = true
		}
	}
	return check
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous 20 locations",
			callback:    commandBMap,
		},
		"explore": {
			name:        "explore",
			description: "Shows available pokemons in specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch specified pokemon(higher level pokemons are harder to catch)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get details about caught pokemons",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows caught pokemons",
			callback:    commndPokedex,
		},
	}
}
