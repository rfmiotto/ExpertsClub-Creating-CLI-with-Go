package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name     string
	commands []Command
}

func CommandInit(name string) *CommandRoot {
	return &CommandRoot{
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandList []Command) error {
	if len(commandList) == 0 {
		return errors.New("no command was registered in the CLI")
	}

	cr.commands = commandList

	if len(os.Args) < 2 {
		cr.showHelp()
		return errors.New("please specify a command")
	}

	userPassedArguments := os.Args[1:]

	userCommand := argumentFilter(userPassedArguments)

	// fmt.Println(userCommand)

	if userCommand.Command == "" {
		cr.showHelp()

		return errors.New("please, specify a valid command")
	}

	if userCommand.Command == "help" {
		cr.showHelp()

		return nil
	}

	for _, command := range cr.commands {
		if userCommand.Command == command.Name() {
			flagSet := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
			command.Register(flagSet)
			flagSet.Parse(os.Args[2:])
			command.Run()
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid command", userCommand.Command)
}

func (cr *CommandRoot) showHelp() {
	fmt.Printf("Usage: %s [COMMAND] [OPTIONS]\n\n", cr.Name)
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tabWriter, "commands:\n\n")
	for _, command := range cr.commands {
		fmt.Fprintf(tabWriter, "\t- %s\t%s\n", command.Name(), command.Help())
	}
	tabWriter.Flush()

	fmt.Fprintf(tabWriter, "\nexamples:\n")
	for _, command := range cr.commands {
		fmt.Fprintf(tabWriter, "\t%s\n", command.Example())
	}
	tabWriter.Flush()
}
