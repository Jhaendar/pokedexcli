package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/Jhaendar/pokedexcli/pokecache"
)

func init() {
	PokemonCaught = make(map[string]Pokemon)
}

var PokemonAPIBaseURL string = "https://pokeapi.co/api/v2/pokemon/%v/"
var PokemonCache = pokecache.NewCache(5 * time.Minute)
var PokemonCaught map[string]Pokemon

func GetPokemon(name_or_id string) (Pokemon, error) {
	url := fmt.Sprintf(PokemonAPIBaseURL, name_or_id)

	var pokemonData Pokemon
	if data, ok := PokemonCache.Get(url); ok {
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonData, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(bodyBytes, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonCache.Add(url, bodyBytes)

	return pokemonData, nil
}

func CatchPokemon(pokemon Pokemon) {
	fmt.Printf("Throwing a pokeball at %v...\n", pokemon.Name)

	//generate random number from 0 to 255
	randomNumber := rand.Intn(256)

	if randomNumber < pokemon.BaseExperience {
		fmt.Println("Oh no! The pokemon broke free!")
		return
	}

	fmt.Printf("Gotcha! %v was caught!\n", pokemon.Name)
	fmt.Printf("You can now inspect this pokemon with `inspect %v\n`", pokemon.Name)
	PokemonCaught[pokemon.Name] = pokemon
}

// New method to get and catch a pokemon
func GetAndCatchPokemon(name_or_id string) error {
	pokemon, err := GetPokemon(name_or_id)
	if err != nil {
		return err
	}

	CatchPokemon(pokemon)

	return nil
}
func GetCaughtPokemon(name string) (Pokemon, bool) {
	pokemon, ok := PokemonCaught[name]
	return pokemon, ok
}
func PrintPokemon(pokemon Pokemon) {
	fmt.Printf("Name: %v\n\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
	fmt.Println()
}

func PrintCaughtPokemon(name string) {
	pokemon, ok := GetCaughtPokemon(name)
	if !ok {
		fmt.Printf("You have not caught %v yet\n", name)
		return
	}

	PrintPokemon(pokemon)
}
