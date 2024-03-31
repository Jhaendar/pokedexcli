// Inspect CLI Command

package cli

import (
	"github.com/Jhaendar/pokedexcli/pokeapi"
)

func init() {
	inspectCLICommand := InspectCLiCommand{
		CLICommandInfo{
			name:        "inspect",
			description: "Inspect a Pokemon",
		}}

	addCLICommand("inspect", inspectCLICommand)
}

type InspectCLiCommand struct {
	CLICommandInfo
}

func (c InspectCLiCommand) Execute(args []string) error {
	pokeapi.PrintCaughtPokemon(args[0])
	return nil
}
