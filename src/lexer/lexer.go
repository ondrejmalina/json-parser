package lexer

type TokenType string

const (
	EOF          TokenType = "EOF"
	INVALID                = "INVALID"
	LEFT_CUR_BR            = "LEFT_CUR_BR"
	RIGHT_CUR_BR           = "RIGHT_CUR_BR"
	LEFT_SQ_BR             = "LEFT_SQ_BR"
	RIGHT_SQ_BR            = "RIGHT_SQ_BR"
)

type Token struct {
	token    TokenType
	value    interface{}
	position int
}

type Lexer struct {
	String    string
	Character byte
	Position  int
}

func (l *Lexer) nextElement() {
	l.Position++
	if l.Position == len(l.String) {
		return
	}
	l.Character = l.String[l.Position]
}

func (l *Lexer) matchToken() Token {
	switch string(l.Character) {
	case "{":
		return Token{LEFT_CUR_BR, l.String[l.Position], l.Position}
	case "}":
		return Token{RIGHT_CUR_BR, l.String[l.Position], l.Position}
	case "[":
		return Token{LEFT_SQ_BR, l.String[l.Position], l.Position}
	case "]":
		return Token{RIGHT_SQ_BR, l.String[l.Position], l.Position}
	default:
		return Token{INVALID, nil, len(l.String)}
	}
}

func (l *Lexer) TokenizeString() []Token {
	var tokens []Token

	for l.Position < len(l.String) {
		token := l.matchToken()
		tokens = append(tokens, token)
		l.nextElement()
	}

	return tokens
}
