package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapb(conf *config) error {
	var requestURL string

	if len(conf.Previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	} else {
		requestURL = conf.Previous
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
