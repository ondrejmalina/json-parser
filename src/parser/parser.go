package parser

import (
	"errors"
	"github.com/ondrejmalina/json-parser/src/lexer"
	"os"
)

// Iterator interface
type Iterator interface {
	hasNext() bool
	getNext() interface{}
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

func (si *ParserIterator) getNext() interface{} {

	currentToken := si.collection[si.position]
	si.position++
	return currentToken
}

func (si *ParserIterator) hasNext() bool {
	if si.position >= len(si.collection) {
		return false
	}
	return true
}

func (p *Parser) ParseJson() {
	if err := p.parseObject(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func (p *Parser) parseObject() error {
	// obj := make(map[string]interface{})

	return errors.New("Error")
}

func (p *Parser) parseValue() {

}
