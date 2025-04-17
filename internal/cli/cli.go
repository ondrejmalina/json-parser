package cli

import (
	"flag"
	"fmt"
	"os"
)

func printHelp() {
	help := `
Usage: json-parser [json input]

Examples:
  - json-parser "{"foo": "bar}"
`
	fmt.Print(help)
}

type UserInput struct {
	Input string
}

func ReadInput() UserInput {
	flag.Usage = printHelp
	flag.Parse()

	inp := flag.Arg(0)

	if len(inp) == 0 {
		printHelp()
		os.Exit(0)
	}

	return UserInput{
		Input: inp,
	}
}
