// HELP CLI COMMAND

package cli

import (
	"fmt"
)

// add helpclicommand to commandsList
func init() {
	helpCLICommand := HelpCLICommand{
		CLICommandInfo{
			name:        "help",
			description: "Display help message",
		}}

	addCLICommand("help", helpCLICommand)
}

type HelpCLICommand struct {
	CLICommandInfo
}

func (h HelpCLICommand) Execute() error {
	fmt.Println(`
    Pokedex CLI
    
    Usage:`)

	for _, command := range CLICommands {
		fmt.Printf("    %s: %s\n", command.GetName(), command.GetDescription())
	}
	return nil
}
