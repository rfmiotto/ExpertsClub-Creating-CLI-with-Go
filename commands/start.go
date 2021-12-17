package commands

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	Helpf   bool
}

const helpTextStart = `Responsible for initializing a simple server`
const helpTextLongStart = `
Initialize a simple server and allow access of all configured services

gocli start --port [serverPort] --version [serverVersion]
- serverPort: int
- serverVersion: string
`
const exampleTextStart = `
gocli start --port 3030 --version 1.0.0
gocli start --version 1.0.0
`

func (cmd *Start) Name() string {
	return "start"
}

func (cmd *Start) Example() string {
	return exampleTextStart
}

func (cmd *Start) Help() string {
	return helpTextStart
}

func (cmd *Start) LongHelp() string {
	return helpTextLongStart
}

func (cmd *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&cmd.Port, "port", 8080, "server port")
	fs.StringVar(&cmd.Version, "version", "", "server version")
	fs.BoolVar(&cmd.Helpf, "help", false, "help command")
}

func (cmd *Start) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Version == "" {
		fmt.Println("[--version] is a required field")
		return
	}

	fmt.Printf("Server %v is running on port %v", cmd.Version, cmd.Port)

	http.ListenAndServe(fmt.Sprintf(":%v", cmd.Port), nil)
}
