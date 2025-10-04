package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *config, _ string) error {
	var url string
	if config.PrevURL == nil {
		url = URL + "?offset=0&limit=20"
	} else {
		url = config.NextURL
	}

	if ok := implementCache(url); ok {
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	cache.Add(url, body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, config); err != nil {
		return fmt.Errorf("Error Unmarhsalling data: %v", err)
	}

	for _, result := range config.Results {
		fmt.Printf("%v\n", result.Name)
	}

	return nil
}
