package message

import (
	"flag"
	"fmt"
)

type Message struct {
	Text  string
	Helpf bool
}

const helpTextStart = `Responsible for printing a simple text`
const helpTextLongStart = `
Print a simple text

gocli message --text [printText]
- printText: string
`
const exampleTextStart = `
gocli message --text hello_world
`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Example() string {
	return exampleTextStart
}

func (cmd *Message) Help() string {
	return helpTextStart
}

func (cmd *Message) LongHelp() string {
	return helpTextLongStart
}

func (cmd *Message) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.Text, "text", "", "text to be printed out")
	fs.BoolVar(&cmd.Helpf, "help", false, "help command")
}

func (cmd *Message) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Text == "" {
		fmt.Println("[--text] is a required field")
		return
	}

	fmt.Printf("The text is: %v", cmd.Text)
}
