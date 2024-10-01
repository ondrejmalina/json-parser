package lexer

type TokenType string

const (
	EOF          TokenType = "EOF"
	INVALID                = "INVALID"
	LEFT_CUR_BR            = "{"
	RIGHT_CUR_BR           = "}"
	LEFT_SQ_BR             = "["
	RIGHT_SQ_BR            = "]"
)

type Token struct {
	Token    TokenType
	Position int
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
		return Token{LEFT_CUR_BR, l.Position}
	case "}":
		return Token{RIGHT_CUR_BR, l.Position}
	case "[":
		return Token{LEFT_SQ_BR, l.Position}
	case "]":
		return Token{RIGHT_SQ_BR, l.Position}
	default:
		return Token{INVALID, l.Position}
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
