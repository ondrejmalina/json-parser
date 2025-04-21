package lexer

import (
	"log"
	"unicode"

	"github.com/ondrejmalina/json-parser/internal/cli"
)

type TokenType string

const (
	EOF          TokenType = "EOF"
	EMPTY                  = "EMPTY"
	INVALID                = "INVALID"
	STRING                 = "STRING"
	DIGIT                  = "DIGIT"
	HIDDEN                 = "HIDDEN"
	NULL                   = "NULL"
	BOOL                   = "BOOL"
	LEFT_CUR_BR            = "{"
	RIGHT_CUR_BR           = "}"
	LEFT_SQ_BR             = "["
	RIGHT_SQ_BR            = "]"
	COLON                  = ":"
	COMMA                  = ","
)

type Lexer struct {
	runes    []rune
	position int
}

type Token struct {
	TokenType TokenType
	Character rune
}

func removeSpaceRunes(runes []rune) []rune {
	noSpaceRunes := make([]rune, 0, len(runes)) // NOTE: preallocates memory for efficiency
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			noSpaceRunes = append(noSpaceRunes, r)
		}
	}
	return noSpaceRunes
}

func CreateLexer(inp cli.UserInput) *Lexer {
	inpRunes := []rune(inp.Input)
	return &Lexer{runes: removeSpaceRunes(inpRunes)}
}

func (l *Lexer) GetToken() Token {

	if l.position >= len(l.runes) {
		return Token{EOF, 0} // 0 is nil in unicode
	}

	r := l.runes[l.position]
	switch {
	case r == '{':
		return Token{LEFT_CUR_BR, r}
	case r == '}':
		return Token{RIGHT_CUR_BR, r}
	case r == ',':
		return Token{COMMA, r}
	case r == ':':
		return Token{COLON, r}
	case r == ':':
		return Token{COMMA, r}
	case r == '"':
		return l.parseString()
	}
	return Token{INVALID, r}
}

func (l *Lexer) NextToken() {
	l.nextRune()
}

func (l *Lexer) nextRune() {
	l.position++
}

func (l *Lexer) parseString() Token {
	l.nextRune()
	for ; l.position < len(l.runes) && l.runes[l.position] != '"'; l.nextRune() {
		log.Print(l.runes[l.position])
		if l.position == len(l.runes)-1 {
			return Token{INVALID, 0}
		}
	}
	l.nextRune() // NOTE: Must shift to another rune to sync with the entire logic
	return Token{STRING, 0}
}
