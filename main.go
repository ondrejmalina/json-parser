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
	lexer := lexer.Lexer{String: file, Position: 0}
	fmt.Println(lexer)
	tokens := lexer.TokenizeString()
	fmt.Println(tokens)

	p := parser.CreateParser(tokens)
	p.ParseJson()
}
