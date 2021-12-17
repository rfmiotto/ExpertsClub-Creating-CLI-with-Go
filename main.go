// To run this script:
// go run main.go --text hello

package main

import (
	"flag"
	"fmt"
	"os"
)

type Example struct {
	Text        string
	FlagSetText string
}

func main() {
	example := Example{}

	flag.StringVar(&example.Text, "text", "default text", "description")

	flag.Parse()

	fmt.Println(example.Text)

	fs := flag.NewFlagSet("add", flag.ExitOnError)

	fs.StringVar(&example.FlagSetText, "text", "default text", "description")

	fs.Parse(os.Args[1:])

	fmt.Println(os.Args)

	fmt.Println(example.FlagSetText)
}
