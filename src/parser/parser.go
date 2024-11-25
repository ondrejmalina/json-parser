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
	case "EMPTY":
		err = errors.New("Empty file is not valid JSON")
	case "{":
		err = p.parseObject()
	}

	return err
}

func (p *parser) parseObject() error {

	// empty json files
	if p.tokens[p.position+1].Token == "}" {
		return nil
	}

	var token lexer.Token
	for true {
		token = p.getNextToken()
		if token.Token != "STRING" {
			return errors.New("JSON key must be a string")
		}

		token = p.getNextToken()
		if token.Token != ":" {
			return errors.New("JSON key must be followed by a colon")
		}

		token = p.getNextToken()
		// if token.Token != "STRING" {
		// 	return errors.New("JSON value must be a string")
		// }

		switch token.Token {
		case "STRING", "BOOL", "NULL", "DIGIT":
			break
		default:
			return errors.New("Invalid value")
		}

		token = p.getNextToken()
		if token.Token == "}" {
			return nil
		}

		if token.Token != "," {
			return errors.New("Missing comma")
		}
	}

	return errors.New("JSON is not valid")
}
