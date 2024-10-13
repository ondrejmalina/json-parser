package lexer

import (
	"testing"
)

func TestMatchToken(t *testing.T) {

	testString := "{}a[]"
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{RIGHT_CUR_BR, 1},
		{INVALID, 2},
		{LEFT_SQ_BR, 3},
		{RIGHT_SQ_BR, 4},
	}
	l := Lexer{testString, 0}

	for l.Position < len(l.String) {
		token := l.matchToken()
		if token != expectedOutput[token.Position] {
			t.Errorf("\nThe input string: %v, position: %v \n"+
				"matches token %v which \n"+
				"does not match expected output %v\n",
				string(l.String[l.Position]),
				l.Position,
				token,
				expectedOutput[l.Position])
		}
		l.nextElement()
	}
}

func TestTokenizeString(t *testing.T) {

	testString := "{[1],}"
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{LEFT_SQ_BR, 1},
		{INVALID, 2},
		{RIGHT_SQ_BR, 3},
		{INVALID, 4},
		{RIGHT_CUR_BR, 5},
	}

	l := Lexer{testString, 0}
	tokens := l.TokenizeString()

	for token := range tokens {
		if tokens[token] != expectedOutput[token] {
			t.Errorf("\nToken %v does not match expected token %v",
				tokens[token],
				expectedOutput[token],
			)
		}
	}
}

func TestTokenizeEmptyString(t *testing.T) {

	testString := ""
	expectedOutput := []Token{
		{EMPTY, 0},
	}

	l := Lexer{testString, 0}
	tokens := l.TokenizeString()

	for token := range tokens {
		if tokens[token] != expectedOutput[token] {
			t.Errorf("\nToken %v does not match expected token %v",
				tokens[token],
				expectedOutput[token],
			)
		}
	}

}
