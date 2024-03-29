// EXIT CLI COMMAND

package cli

import (
	"fmt"
	"os"
)

func init() {
	exitCLICommand := ExitCLICommand{
		CLICommandInfo{
			name:        "exit",
			description: "Exit Pokedex CLI",
		}}

	addCLICommand("exit", exitCLICommand)
}

type ExitCLICommand struct {
	CLICommandInfo
}

func (e ExitCLICommand) Execute() error {
	fmt.Println("Exiting Pokedex CLI...")
	os.Exit(0)
	return nil
}
