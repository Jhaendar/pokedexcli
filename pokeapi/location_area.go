package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Jhaendar/pokedexcli/pokecache"
)

type LocationAreaIndex struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Resource `json:"results"`
}

var LocationAreaIndexURL = "https://pokeapi.co/api/v2/location-area/"

var LocationAreaIndexCache = pokecache.NewCache(5 * time.Minute)

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
func PrintResourceNames(resources []Resource) {
	for _, resource := range resources {
		fmt.Println(resource.Name)
	}
}
