// Explore CLI Command

package cli

import (
	"errors"

	"github.com/Jhaendar/pokedexcli/pokeapi"
)

func init() {
	exploreCLICommand := ExploreCLICommand{
		CLICommandInfo{
			name:        "explore",
			description: "Explore the Pokedex",
		}}

	addCLICommand("explore", exploreCLICommand)
}

type ExploreCLICommand struct {
	CLICommandInfo
}

func (e ExploreCLICommand) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments provided")
	}

	locationArea, err := pokeapi.GetLocationArea(args[0])
	if err != nil {
		return errors.New("error getting location area")
	}

	pokeapi.PrintPokemonEncounters(locationArea)

	return nil
}
