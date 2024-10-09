package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/ondrejmalina/json-parser/src/lexer"
)

func CreateParser(tokens []lexer.Token) parser {
	return parser{tokens: tokens, position: 0}
}

type parser struct {
	tokens   []lexer.Token
	position int
}

func (p *parser) getNextToken() lexer.Token {
	p.position++
	return p.tokens[p.position]
}

func (p *parser) ParseJson() {
	err := p.parseValue()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Json file parsed successfully")
	os.Exit(0)
}

func (p *parser) parseValue() error {

	var err error
	err = nil

	switch p.tokens[p.position].Token {
	case "{":
		err = p.parseObject()
	}

	return err
}

func (p *parser) parseObject() error {
	token := p.getNextToken()
	for true {
		if token.Token == "}" {
			return nil
		}
	}
	return errors.New("JSON is not valid")
}
