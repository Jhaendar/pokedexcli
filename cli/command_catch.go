// Catch CLI command

package cli

import "github.com/Jhaendar/pokedexcli/pokeapi"

func init() {
	catchCLICommand := CatchCLiCommand{
		CLICommandInfo{
			name:        "catch",
			description: "Catch a Pokemon",
		}}

	addCLICommand("catch", catchCLICommand)
}

type CatchCLiCommand struct {
	CLICommandInfo
}

func (c CatchCLiCommand) Execute(args []string) error {
	err := pokeapi.GetAndCatchPokemon(args[0])
	if err != nil {
		return err
	}
	return nil
}
