package cli

var CLICommands map[string]CLICommand

func addCLICommand(name string, command CLICommand) {
	if CLICommands == nil {
		CLICommands = make(map[string]CLICommand)
	}
	CLICommands[name] = command
}

type CLICommand interface {
	Execute() error
	GetName() string
	GetDescription() string
}

type CLICommandInfo struct {
	name        string
	description string
}

func (c CLICommandInfo) GetName() string {
	return c.name
}

func (c CLICommandInfo) GetDescription() string {
	return c.description
}
