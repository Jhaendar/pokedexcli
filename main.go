package main

import (
	"fmt"
	"strings"

	"github.com/Jhaendar/pokedexcli/cli"
)

func main() {
	commands := cli.CLICommands

	for {
		var input string

		fmt.Print("pokedex > ")
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Error reading command")
		}

		command, args := processInput(input)

		cliCommand, commandExists := commands[command]
		if !commandExists {
			fmt.Println("Command not found")
			continue
		}

		commandError := cliCommand.Execute(args)
		if commandError != nil {
			fmt.Println("Error executing command: ", commandError)
		}

	}
}

func processInput(input string) (string, []string) {
	input_slice := strings.Split(input, " ")

	return input_slice[0], input_slice[1:]
}
