package parser

import (
	"fmt"

	"github.com/ondrejmalina/json-parser/internal/lexer"
)

func CreateParser(l lexer.Lexer) Parser {
	token := l.TokenizeRune()
	l.NextElement()
	nToken := l.TokenizeRune()
	return Parser{lexer: l, token: token, nextToken: nToken}
}

type Parser struct {
	lexer     lexer.Lexer
	token     lexer.Token
	nextToken lexer.Token
}

func (p *Parser) ParseJson() error {
	err := p.parseJson()
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) getNextToken() {
	p.token = p.nextToken
	p.lexer.NextElement()
	p.nextToken = p.lexer.TokenizeRune()
}

func (p *Parser) parseJson() error {

	// check valid start of json
	if p.token.TokenType != lexer.LEFT_CUR_BR {
		return fmt.Errorf("JSON starts with invalid character '%c'", p.token.Character)
	}

	// check valid json parsing
	if err := p.parseObject(); err != nil {
		return err
	}

	// check no characters after json closure
	if p.nextToken.TokenType != lexer.EOF {
		return fmt.Errorf("JSON has invalid closure '%c'", p.nextToken.Character)
	}

	return nil
}

func (p *Parser) parseObject() error {

	for true {
		p.getNextToken()

		switch p.token.TokenType {
		case lexer.RIGHT_CUR_BR:
			return nil
		default:
			return fmt.Errorf("Invalid value %c", p.token.Character)
		}
	}
	return nil
}
