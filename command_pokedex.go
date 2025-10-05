package main

import "fmt"

func commndPokedex(_ *config, _ string) error {
	if len(MyPokemons) == 0 {
		return fmt.Errorf("your pokedex is empty(just like your brain)")
	}

	fmt.Printf("Your Pokedex:\n")
	for name := range MyPokemons {
		fmt.Printf("	- %v\n", name)
	}
	return nil
}
