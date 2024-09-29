package parser

import (
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
	p.parseValue()
}

func (p *parser) parseObject() {
	token := p.getNextToken()
	for true {
		if token.Token == "}" {
			os.Exit(0)
		}
	}
	os.Exit(1)
}

func (p *parser) parseValue() {
	switch p.tokens[p.position].Token {
	case "{":
		p.parseObject()
	}
}
