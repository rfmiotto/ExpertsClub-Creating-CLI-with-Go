package internal

import (
	"regexp"
)

type UserCommand struct {
	Command   string
	Arguments []string // arguments are flags passed to the command.
}

// [executable command --someFlag --someValue]
// the commandList will be: [command --someFlag --someValue]
// example:
// go run cmd/Go_CLI_example/main.go start --port 3333 --version 1.0
// the commandList will be:  [start --port 3333 --version]
func argumentFilter(commandList []string) UserCommand {
	regexpValidator := regexp.MustCompile("(?m)-")

	// reads the first item in the slice and checks if it is a flag or a command.
	// If it is a command, the `commandSet` below will be true. This means that
	// the command is already set and everything that comes after that are flags.
	commandSet := false

	// fmt.Println(commandList)

	var commandDefinition UserCommand

	for _, argument := range commandList {
		flagMatches := regexpValidator.MatchString(argument)

		if !flagMatches && !commandSet {
			commandDefinition.Command = argument
			commandSet = true
		} else if flagMatches {
			commandDefinition.Arguments = append(commandDefinition.Arguments, argument)
		}
	}

	return commandDefinition
}
