package parser

import (
	"fmt"

	"github.com/ondrejmalina/json-parser/internal/lexer"
)

func CreateParser(l lexer.Lexer) Parser {
	return Parser{l}
}

type Parser struct {
	lexer lexer.Lexer
}

func (p *Parser) ParseJson() error {
	err := p.parseValue()
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) parseValue() error {

	token := p.lexer.TokenizeRune()

	switch token.TokenType {
	case lexer.LEFT_CUR_BR:
		return p.parseObject()
	}

	return fmt.Errorf("JSON starts with invalid character '%c'", token.Character)
}

func (p *Parser) parseObject() error {

	for true {
		p.lexer.NextElement()
		token := p.lexer.TokenizeRune()

		switch token.TokenType {
		case lexer.RIGHT_CUR_BR:
			return nil
		case lexer.EOF:
			return fmt.Errorf("Reached EOF")
		default:
			return fmt.Errorf("Invalid value %c", token.Character)
		}
	}
	return nil
}
