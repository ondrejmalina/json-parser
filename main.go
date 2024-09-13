package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type TokenType int

const (
	EOF = iota // initator, starts at 0, adds 1 to each successive non-empty, non-comment line
	INVALID
	RIGHT_CUR_BR
	LEFT_CUR_BR
	LEFT_SQ_BR
	RIGHT_SQ_BR
)

func Read() []byte {
	file := flag.Args()
	if len(file) == 0 {
		log.Fatal("No file input")
	}

	text, err := os.ReadFile(file[0])
	if err != nil {
		log.Fatal(err)
	}

	return text
}

type Token struct {
	token    TokenType
	value    interface{}
	position int
}

type Lexer struct {
	str       string
	character byte
	position  int
}

func (l *Lexer) nextElement() {
	l.position++
}

func (l *Lexer) matchToken() Token {
	switch string(l.character) {
	case "{":
		return Token{LEFT_CUR_BR, l.str[l.position], l.position}
	case "}":
		return Token{RIGHT_CUR_BR, l.str[l.position], l.position}
	case "[":
		return Token{LEFT_SQ_BR, l.str[l.position], l.position}
	case "]":
		return Token{RIGHT_SQ_BR, l.str[l.position], l.position}
	default:
		return Token{INVALID, nil, len(l.str)}
	}
}

func (l *Lexer) TokenizeString() []Token {
	var tokens []Token
	lenStr := len(l.str)

	for i := 0; i < lenStr; i++ {
		token := l.matchToken()
		tokens = append(tokens, token)
		l.nextElement()
	}

	return tokens
}

func main() {

	flag.Parse()

	f := string(Read())
	l := Lexer{f, f[0], 0}

	fmt.Println(l.TokenizeString())

}
