package lexer

type TokenType string

const (
	EOF          TokenType = "EOF"
	EMPTY                  = "EMPTY"
	INVALID                = "INVALID"
	STRING                 = "STRING"
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
	switch l.Runes[l.Position] {
	case '{':
		return Token{LEFT_CUR_BR, l.Position}
	case '}':
		return Token{RIGHT_CUR_BR, l.Position}
	case '[':
		return Token{LEFT_SQ_BR, l.Position}
	case ']':
		return Token{RIGHT_SQ_BR, l.Position}
	case ',':
		return Token{COMMA, l.Position}
	case ':':
		return Token{COLON, l.Position}
	case '\n', ' ', '\t':
		return Token{HIDDEN, l.Position}
	case '"':
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
