package main

import (
	"fmt"

	"github.com/rfmiotto/Go_CLI_example/commands"
	"github.com/rfmiotto/Go_CLI_example/commands/message"
	"github.com/rfmiotto/Go_CLI_example/internal"
)

var commandList = []internal.Command{
	new(commands.Start),
	new(message.Message),
}

func main() {
	err := internal.CommandInit("gocli").Start(commandList)

	if err != nil {
		fmt.Println(err.Error())
	}
}
