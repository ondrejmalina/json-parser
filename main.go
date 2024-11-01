package main

import (
	"flag"

	"github.com/ondrejmalina/json-parser/src/file"
	"github.com/ondrejmalina/json-parser/src/lexer"
	"github.com/ondrejmalina/json-parser/src/parser"
)

func main() {
	flag.Parse()

	file := file.Read()
	lexer := lexer.Lexer{Runes: file, Position: 0}
	tokens := lexer.TokenizeString()

	p := parser.CreateParser(tokens)
	p.ParseJson()
}
