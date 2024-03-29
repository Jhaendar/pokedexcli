package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaIndex struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Resource `json:"results"`
}

var LocationAreaIndexURL = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreaIndex(url string) (LocationAreaIndex, error) {
	if url == "" {
		url = LocationAreaIndexURL
	}
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaIndex{}, err
	}

	defer resp.Body.Close()

	var data LocationAreaIndex
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return LocationAreaIndex{}, err

	}

	return data, nil
}
func PrintResourceNames(resources []Resource) {
	for _, resource := range resources {
		fmt.Println(resource.Name)
	}
}
