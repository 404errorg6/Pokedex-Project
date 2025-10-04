package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandBMap(config *config, _ string) error {
	if config.PrevURL == nil {
		fmt.Printf("you're on first page\n")
		return nil
	}

	if ok := implementCache(*config.PrevURL); ok {
		return nil
	}

	res, err := http.Get(*config.PrevURL)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Error occured with status code: %d and\nbody: %v", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, config); err != nil {
		return fmt.Errorf("Error while Unmarshalling: %v", err)
	}

	for _, result := range config.Results {
		fmt.Printf("%v\n", result.Name)
	}

	return nil
}
