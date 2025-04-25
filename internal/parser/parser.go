package parser

import (
	"errors"
	"fmt"

	"github.com/ondrejmalina/json-parser/internal/lexer"
)

func CreateParser(l lexer.Lexer) Parser {
	token := l.GetToken()
	l.NextToken()
	nToken := l.GetToken()
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
	p.lexer.NextToken()
	p.nextToken = p.lexer.GetToken()
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

	// End for empty objects
	if p.nextToken.TokenType == lexer.RIGHT_CUR_BR {
		p.getNextToken() // NOTE: For JSON closure check
		return nil
	}

	for true {
		// key
		p.getNextToken()
		if p.token.TokenType != lexer.STRING {
			return fmt.Errorf("JSON key must be a string")
		}

		// colon
		p.getNextToken()
		if p.token.TokenType != lexer.COLON {
			return fmt.Errorf("Missing colon after key")
		}

		// value
		p.getNextToken()
		switch p.token.TokenType {
		case lexer.STRING, lexer.DIGIT, lexer.BOOL, lexer.NULL:
			break
		case lexer.LEFT_CUR_BR:
			p.parseObject()
		case lexer.LEFT_SQ_BR:
			p.parseArray()
		default:
			return fmt.Errorf("Invalid JSON value")
		}

		// end of object or comma
		p.getNextToken()
		switch {
		case p.token.TokenType == lexer.RIGHT_CUR_BR:
			return nil
		case p.token.TokenType != ",":
			return errors.New("Missing comma")
		}
	}
	return nil
}

func (p *Parser) parseArray() error {

	// end for empty arrays
	if p.nextToken.TokenType == lexer.RIGHT_SQ_BR {
		p.getNextToken()
		return nil
	}

	for true {

		// Values
		p.getNextToken()
		switch p.token.TokenType {
		case lexer.STRING, lexer.DIGIT, lexer.BOOL, lexer.NULL:
			break
		default:
			return fmt.Errorf("Invalid array value")
		}

		// End of array or comma
		p.getNextToken()
		switch {
		case p.token.TokenType == lexer.RIGHT_SQ_BR:
			return nil
		case p.token.TokenType != ",":
			return errors.New("Missing comma")
		}
	}
	return nil
}
