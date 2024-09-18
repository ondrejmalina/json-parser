package main

import (
	"flag"
	"fmt"
	"github.com/ondrejmalina/json-parser/src/file"
	"github.com/ondrejmalina/json-parser/src/lexer"
)

var Parser struct {
	tokens []lexer.Token
}

func main() {
	flag.Parse()

	file := string(file.Read())
	lexer := lexer.Lexer{String: file, Character: file[0], Position: 0}

	fmt.Println(lexer.TokenizeString())
}
