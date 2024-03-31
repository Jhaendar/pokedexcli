package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Jhaendar/pokedexcli/pokecache"
)

const LocationAreaIndexURL = "https://pokeapi.co/api/v2/location-area/"

var LocationAreaIndexCache = pokecache.NewCache(5 * time.Minute)
var LocationAreaCache = pokecache.NewCache(5 * time.Minute)

func GetLocationAreaIndex(url string) (LocationAreaIndex, error) {
	if url == "" {
		url = LocationAreaIndexURL
	}

	var locationAreaIndexData LocationAreaIndex

	if data, ok := LocationAreaIndexCache.Get(url); ok {
		err := json.Unmarshal(data, &locationAreaIndexData)
		if err != nil {
			return LocationAreaIndex{}, err
		}
		return locationAreaIndexData, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaIndex{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaIndex{}, err
	}

	LocationAreaIndexCache.Add(url, bodyBytes)

	err = json.Unmarshal(bodyBytes, &locationAreaIndexData)
	if err != nil {
		return LocationAreaIndex{}, err
	}

	return locationAreaIndexData, nil
}

func GetLocationArea(name_or_id string) (LocationArea, error) {
	var locationAreaData LocationArea

	if data, ok := LocationAreaCache.Get(name_or_id); ok {
		err := json.Unmarshal(data, &locationAreaData)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaData, nil
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/" + name_or_id)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	LocationAreaCache.Add(name_or_id, bodyBytes)
	err = json.Unmarshal(bodyBytes, &locationAreaData)
	if err != nil {
		return LocationArea{}, err
	}

	return locationAreaData, nil
}

func PrintResourceNames(resources []Resource) {
	for _, resource := range resources {
		fmt.Println(resource.Name)
	}
}

func PrintPokemonEncounters(locationArea LocationArea) {
	fmt.Printf("Exploring %v\n", locationArea.Name)
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
}
