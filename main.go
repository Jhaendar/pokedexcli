package main

import (
	"fmt"

	"github.com/Jhaendar/pokedexcli/cli"
)

func main() {
	commands := cli.CLICommands

	for {
		var command string

		fmt.Print("pokedex > ")
		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Println("Error reading command")
		}

		cliCommand, commandExists := commands[command]
		if !commandExists {
			fmt.Println("Command not found")
			continue
		}

		commandError := cliCommand.Execute()
		if commandError != nil {
			fmt.Println("Error executing command: ", commandError)
		}

	}
}
