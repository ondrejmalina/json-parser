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

func (l *Lexer) TokenizeRune() Token {

	if l.position >= len(l.runes) {
		return Token{EOF, 0} // 0 is nil in unicode
	}

	r := l.runes[l.position]
	switch {
	case r == '{':
		return Token{LEFT_CUR_BR, r}
	case r == '}':
		return Token{RIGHT_CUR_BR, r}
	}
	return Token{INVALID, r}
}

func (l *Lexer) NextElement() {
	l.position++
}

// func (l *Lexer) matchToken() Token {
//
// special_runes := rangetable.New('n', 't', 'f')

// rune := l.Runes[l.CurrentPosition]
// switch {
// case rune == '{':
// 	return Token{LEFT_CUR_BR, l.CurrentPosition}
// case rune == '}':
// 	return Token{RIGHT_CUR_BR, l.CurrentPosition}
// case rune == '[':
// 	return Token{LEFT_SQ_BR, l.Position}
// case rune == ']':
// 	return Token{RIGHT_SQ_BR, l.Position}
// case rune == ',':
// 	return Token{COMMA, l.Position}
// case rune == ':':
// 	return Token{COLON, l.Position}
// case unicode.IsSpace(rune) == true:
// 	return Token{HIDDEN, l.Position}
// // TODO: Add support for negative numbers
// case unicode.IsDigit(rune) == true:
// 	return l.parseDigit()
// case rune == '"':
// 	return l.parseString()
// case unicode.In(rune, special_runes) == true:
// 	return l.parseSpecialFactory()
// default:
// 	return Token{INVALID, l.CurrentPosition}
// }
// }

// func (l *Lexer) TokenizeString() []Token {
// 	// NOTE: Where to add condition to return error
// 	// if list of tokens has INVALID token?
// 	// so far parser does it
//
// 	var tokens []Token
//
// 	if len(l.Runes) == 0 {
// 		token := Token{EMPTY, 0}
// 		tokens = append(tokens, token)
// 		return tokens
// 	}
//
// 	for l.CurrentPosition < len(l.Runes) {
// 		token := l.matchToken()
// 		switch {
// 		case token.Token == INVALID:
// 			tokens = append(tokens, token)
// 			return tokens

// TODO: Using nextElement differently for digits,

// bool and null is ugly, improve
// case token.Token == HIDDEN:
// 	l.nextElement()
// case token.Token == DIGIT:
// 	tokens = append(tokens, token)
// case token.Token == BOOL:
// 	tokens = append(tokens, token)
// case token.Token == NULL:
// 	tokens = append(tokens, token)
// 		default:
// 			tokens = append(tokens, token)
// 			l.nextElement()
// 		}
// 	}
// 	return tokens
// }
//
// // TODO: Rewrite this to return error
// func (l *Lexer) nextElement() {
// 	l.CurrentPosition++
// 	if l.CurrentPosition == len(l.Runes) {
// 		return
// 	}
// }
//
// func (l *Lexer) matchToken() Token {
//
// special_runes := rangetable.New('n', 't', 'f')

// rune := l.Runes[l.CurrentPosition]
// switch {
// case rune == '{':
// 	return Token{LEFT_CUR_BR, l.CurrentPosition}
// case rune == '}':
// 	return Token{RIGHT_CUR_BR, l.CurrentPosition}
// case rune == '[':
// 	return Token{LEFT_SQ_BR, l.Position}
// case rune == ']':
// 	return Token{RIGHT_SQ_BR, l.Position}
// case rune == ',':
// 	return Token{COMMA, l.Position}
// case rune == ':':
// 	return Token{COLON, l.Position}
// case unicode.IsSpace(rune) == true:
// 	return Token{HIDDEN, l.Position}
// // TODO: Add support for negative numbers
// case unicode.IsDigit(rune) == true:
// 	return l.parseDigit()
// case rune == '"':
// 	return l.parseString()
// case unicode.In(rune, special_runes) == true:
// 	return l.parseSpecialFactory()
// default:
// 	return Token{INVALID, l.CurrentPosition}
// }
// }

// func (l *Lexer) parseSpecialFactory() Token {
// 	switch l.Runes[l.Position] {
// 	case 'n':
// 		return l.parseSpecial("null", 4, Token{NULL, l.Position})
// 	case 't':
// 		return l.parseSpecial("true", 4, Token{BOOL, l.Position})
// 	case 'f':
// 		return l.parseSpecial("false", 5, Token{BOOL, l.Position})
// 	}
// 	return Token{INVALID, l.Position}
// }
//
// func (l *Lexer) parseSpecial(exp_string string, len_exp_string int, return_token Token) Token {
// 	// TODO: Where to set position if token is invalid?
// 	// I think lexer should stop completely and let the parser fail
// 	// Lexer must be part of parser

// 	startingPosition := l.Position
// 	collected_string := ""
//
// 	for i := 0; i < len_exp_string; i++ {
// 		if l.Position == len(l.Runes) {
// 			return Token{INVALID, startingPosition}
// 		}
// 		collected_string = collected_string + string(l.Runes[l.Position])
// 		l.nextElement()
// 	}
//
// 	if collected_string == exp_string {
// 		return return_token
// 	}
//
// 	return Token{INVALID, startingPosition}
// }
//
// func (l *Lexer) parseString() Token {

// 	// TODO: How to deal with empty strings?

// 	startingPosition := l.Position
// 	for l.nextElement(); l.Runes[l.Position] != '"'; {
// 		l.nextElement()
// 		if l.Position == len(l.Runes) {
// 			return Token{INVALID, startingPosition}
// 		}
// 	}
//
// 	return Token{STRING, startingPosition}
// }
//
// func (l *Lexer) parseDigit() Token {
// 	startingPosition := l.Position
// 	for unicode.IsDigit(l.Runes[l.Position]) {
// 		if l.Position == len(l.Runes)-1 {
// 			return Token{DIGIT, startingPosition}
// 		}
// 		l.nextElement()
// 	}
//
// 	return Token{DIGIT, startingPosition}
// }
