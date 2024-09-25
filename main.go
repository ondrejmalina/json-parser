package main

import (
	"flag"
	"fmt"
	"github.com/ondrejmalina/json-parser/src/file"
	"github.com/ondrejmalina/json-parser/src/lexer"
	"github.com/ondrejmalina/json-parser/src/parser"
)

func main() {
	flag.Parse()

	file := string(file.Read())
	lexer := lexer.Lexer{String: file, Character: file[0], Position: 0}
	tokens := lexer.TokenizeString()
	fmt.Println(tokens)

	p := parser.Parser{Tokens: tokens}
	p.ParseJson()
}
