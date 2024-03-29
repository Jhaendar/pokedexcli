package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jhaendar/pokedexcli/cli"
)

func main() {
	commands := cli.CLICommands

	for {
		fmt.Print("pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading command:", err)
		}

		command, args := processInput(input)

		cliCommand, commandExists := commands[command]
		if !commandExists {
			fmt.Println("Command not found: ", command, input)
			continue
		}

		commandError := cliCommand.Execute(args)
		if commandError != nil {
			fmt.Println("Error executing command: ", commandError)
		}
	}
}

func processInput(input string) (string, []string) {
	input = strings.TrimSpace(input)
	input_slice := strings.Split(input, " ")

	return input_slice[0], input_slice[1:]
}
