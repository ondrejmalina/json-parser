package cmd

import (
	"fmt"

	"github.com/ondrejmalina/json-parser/internal/cli"
)

func Execute() {

	userInput := cli.ReadInput()
	fmt.Print(userInput)
}
