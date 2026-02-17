package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.prevLocationsURL = response.Previous

	for _, l := range response.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	response, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.prevLocationsURL = response.Previous

	for _, l := range response.Results {
		fmt.Println(l.Name)
	}
	return nil
}
