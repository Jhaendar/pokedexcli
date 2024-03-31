// Pokedex CLI Command

package cli

import "github.com/Jhaendar/pokedexcli/pokeapi"

func init() {
	pokedexCLICommand := PokedexCLiCommand{
		CLICommandInfo{
			name:        "pokedex",
			description: "List all caught Pokemon",
		}}

	addCLICommand("pokedex", pokedexCLICommand)
}

type PokedexCLiCommand struct {
	CLICommandInfo
}

func (c PokedexCLiCommand) Execute(args []string) error {
	pokeapi.PrintAllCaughtPokemon()
	return nil
}
