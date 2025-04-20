package cmd

import (
	"github.com/ondrejmalina/json-parser/internal/cli"
	"github.com/ondrejmalina/json-parser/internal/lexer"
	"github.com/ondrejmalina/json-parser/internal/parser"
)

func Execute() error {

	// userInput
	userInput := cli.ReadInput()

	// lexer
	lexer := lexer.CreateLexer(userInput)

	// parser
	parser := parser.CreateParser(*lexer)
	err := parser.ParseJson()
	if err != nil {
		return err
	}

	return nil
}
