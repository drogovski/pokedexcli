package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const defaultURL = "https://pokeapi.co/api/v2/location-area/"

type pokedexResponse struct {
	Next     string
	Previous string
	Results  []map[string]interface{}
}

func commandMap(conf *config) error {
	var requestURL string

	if len(conf.Next) != 0 {
		requestURL = conf.Next
	} else {
		requestURL = defaultURL
	}

	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var result pokedexResponse
	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	newConf := config{Next: result.Next, Previous: result.Previous}
	*conf = newConf
	showLocations(result.Results)
	return nil
}

func showLocations(locations []map[string]interface{}) {
	for _, location := range locations {
		fmt.Println(location["name"])
	}
}
