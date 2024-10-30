package lexer

type TokenType string

const (
	EOF          TokenType = "EOF"
	EMPTY                  = "EMPTY"
	INVALID                = "INVALID"
	STRING                 = "STRING"
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
	String   string
	Position int
}

func (l *Lexer) nextElement() {
	l.Position++
	if l.Position == len(l.String) {
		return
	}
}

func (l *Lexer) matchToken() Token {
	switch string(l.String[l.Position]) {
	case "{":
		return Token{LEFT_CUR_BR, l.Position}
	case "}":
		return Token{RIGHT_CUR_BR, l.Position}
	case "[":
		return Token{LEFT_SQ_BR, l.Position}
	case "]":
		return Token{RIGHT_SQ_BR, l.Position}
	case ",":
		return Token{COMMA, l.Position}
	case ":":
		return Token{COLON, l.Position}
	case `"`:
		return l.parseString()
	default:
		return Token{INVALID, l.Position}
	}
}

func (l *Lexer) parseString() Token {
	startingPosition := l.Position
	for l.nextElement(); string(l.String[l.Position]) != `"`; {
		l.nextElement()
		if l.Position == len(l.String) {
			return Token{INVALID, startingPosition}
		}
	}

	return Token{STRING, startingPosition}
}

func (l *Lexer) TokenizeString() []Token {
	var tokens []Token

	if len(l.String) == 0 {
		token := Token{EMPTY, 0}
		tokens = append(tokens, token)
		return tokens
	}

	for l.Position < len(l.String) {
		token := l.matchToken()
		tokens = append(tokens, token)
		l.nextElement()
	}

	return tokens
}
