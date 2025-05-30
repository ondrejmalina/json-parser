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
	NUMBER                 = "NUMBER"
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
		return Token{EOF, 0}
	}

	r := l.runes[l.position]

	// NOTE: skip whitespaces
	for l.position < len(l.runes) && unicode.IsSpace(r) {
		l.position++
	}

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
	case r == '-':
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

	digitFound := false

	// Optional leading minus
	if l.runes[l.position] == '-' {
		l.nextRune()
		if !unicode.IsDigit(l.runes[l.position]) {
			// Minus must be followed by at least one digit
			return Token{INVALID, 0}
		}
		for unicode.IsDigit(l.runes[l.position]) {
			digitFound = true
			l.nextRune()
		}
	}

	// Integer part
	for unicode.IsDigit(l.runes[l.position]) {
		digitFound = true
		l.nextRune()
	}

	// Fractional part
	if l.runes[l.position] == '.' {
		l.nextRune()
		if !unicode.IsDigit(l.runes[l.position]) {
			// A dot must be followed by at least one digit
			return Token{INVALID, 0}
		}
		for unicode.IsDigit(l.runes[l.position]) {
			digitFound = true
			l.nextRune()
		}
	}

	// Exponent part
	if r := l.runes[l.position]; r == 'e' || r == 'E' {
		l.nextRune()
		if r := l.runes[l.position]; r == '+' || r == '-' {
			l.nextRune()
		}
		if !unicode.IsDigit(l.runes[l.position]) {
			// Exponent must have at least one digit
			return Token{INVALID, 0}
		}
		for unicode.IsDigit(l.runes[l.position]) {
			digitFound = true
			l.nextRune()
		}
	}

	if !digitFound {
		return Token{INVALID, 0}
	}

	l.previousRune()
	return Token{NUMBER, 0}
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
