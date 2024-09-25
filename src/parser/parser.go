package parser

import (
	"errors"
	"github.com/ondrejmalina/json-parser/src/lexer"
	"os"
)

// Iterator interface
type Iterator interface {
	next() (interface{}, bool)
}

type Collection interface {
	CreateIterator() Iterator
}

// Parse Iterator
type Parser struct {
	Tokens []lexer.Token
}

func (p *Parser) CreateIterator() Iterator {
	return &ParserIterator{
		collection: p.Tokens,
		position:   0,
	}
}

type ParserIterator struct {
	collection []lexer.Token
	position   int
}

func (si *ParserIterator) next() (interface{}, bool) {
	if si.position >= len(si.collection) {
		return nil, false
	}
	si.position++
	return si.collection[si.position], true
}

func (p *Parser) ParseJson() {
	if err := p.parseToken(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func (p *Parser) parseToken() error {
	// obj := make(map[string]interface{})

	return errors.New("Error")

}
