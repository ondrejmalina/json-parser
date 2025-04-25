package lexer

import (
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
	noSpaceRunes := make([]rune, 0, len(runes))
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
	case r == '[':
		return Token{LEFT_SQ_BR, r}
	case r == ']':
		return Token{RIGHT_SQ_BR, r}
	case r == ',':
		return Token{COMMA, r}
	case r == ':':
		return Token{COLON, r}
	case r == '"':
		return l.parseString()
	case unicode.IsDigit(r):
		return l.parseDigit()
	case r == 't' || r == 'f':
		return l.parseBool()
	case r == 'n':
		return l.parseNull()
	}
	return Token{INVALID, r}
}

func (l *Lexer) NextToken() {
	l.nextRune()
}

func (l *Lexer) nextRune() {
	l.position++
}

func (l *Lexer) previousRune() {
	l.position--
}

func (l *Lexer) parseString() Token {
	l.nextRune()
	for ; l.position < len(l.runes) && l.runes[l.position] != '"'; l.nextRune() {
		if l.position == len(l.runes)-1 {
			return Token{INVALID, 0}
		}
	}
	return Token{STRING, 0}
}

func (l *Lexer) parseDigit() Token {
	for ; l.position < len(l.runes) && unicode.IsDigit(l.runes[l.position]); l.nextRune() {
		if l.position == len(l.runes)-1 {
			return Token{INVALID, 0}
		}
	}

	// TODO: Ugly, introduce something like next rune to handle better?
	l.previousRune() // NOTE: ends on rune not being digit, therefore must shift back
	return Token{DIGIT, 0}
}

func (l *Lexer) parseBool() Token {

	positive := "true"
	negative := "false"
	parsedBool := ""

	for ; l.position < len(l.runes); l.nextRune() {
		if l.position == len(l.runes)-1 {
			return Token{INVALID, 0}
		}

		if len(parsedBool) > len(negative) {
			return Token{INVALID, 0}
		}

		parsedBool += string(l.runes[l.position])
		if parsedBool == positive || parsedBool == negative {
			break
		}
	}
	return Token{BOOL, 0}
}

func (l *Lexer) parseNull() Token {

	parsedNull := ""
	strNull := "null"

	for ; l.position < len(l.runes); l.nextRune() {
		if l.position == len(l.runes)-1 {
			return Token{INVALID, 0}
		}

		if len(parsedNull) > len(strNull) {
			return Token{INVALID, 0}
		}

		parsedNull += string(l.runes[l.position])
		if parsedNull == strNull {
			break
		}
	}
	return Token{NULL, 0}
}
