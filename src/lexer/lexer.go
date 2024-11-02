package lexer

import "unicode"

type TokenType string

const (
	EOF          TokenType = "EOF"
	EMPTY                  = "EMPTY"
	INVALID                = "INVALID"
	STRING                 = "STRING"
	DIGIT                  = "DIGIT"
	HIDDEN                 = "HIDDEN"
	LEFT_CUR_BR            = "{"
	RIGHT_CUR_BR           = "}"
	LEFT_SQ_BR             = "["
	RIGHT_SQ_BR            = "]"
	COLON                  = ":"
	COMMA                  = ","
)

type Token struct {
	Token    TokenType
	Position int
}

type Lexer struct {
	Runes    []rune
	Position int
}

func (l *Lexer) nextElement() {
	l.Position++
	if l.Position == len(l.Runes) {
		return
	}
}

func (l *Lexer) matchToken() Token {
	rune := l.Runes[l.Position]
	switch {
	case rune == '{':
		return Token{LEFT_CUR_BR, l.Position}
	case rune == '}':
		return Token{RIGHT_CUR_BR, l.Position}
	case rune == '[':
		return Token{LEFT_SQ_BR, l.Position}
	case rune == ']':
		return Token{RIGHT_SQ_BR, l.Position}
	case rune == ',':
		return Token{COMMA, l.Position}
	case rune == ':':
		return Token{COLON, l.Position}
	case unicode.IsSpace(rune) == true:
		return Token{HIDDEN, l.Position}
	case rune == '"':
		return l.parseString()
	default:
		return Token{INVALID, l.Position}
	}
}

func (l *Lexer) parseString() Token {
	// TODO: How to deal with empty strings?
	startingPosition := l.Position
	for l.nextElement(); string(l.Runes[l.Position]) != `"`; {
		l.nextElement()
		if l.Position == len(l.Runes) {
			return Token{INVALID, startingPosition}
		}
	}

	return Token{STRING, startingPosition}
}

func (l *Lexer) parseDigit() Token {
	startingPosition := l.Position
	for unicode.IsDigit(l.Runes[l.Position]) {
		if l.Position == len(l.Runes)-1 {
			return Token{DIGIT, startingPosition}
		}
		l.nextElement()
	}

	return Token{DIGIT, startingPosition}
}

func (l *Lexer) TokenizeString() []Token {
	var tokens []Token

	if len(l.Runes) == 0 {
		token := Token{EMPTY, 0}
		tokens = append(tokens, token)
		return tokens
	}

	for l.Position < len(l.Runes) {
		token := l.matchToken()
		if token.Token != HIDDEN {
			tokens = append(tokens, token)
		}
		l.nextElement()
	}

	return tokens
}
