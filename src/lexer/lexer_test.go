package lexer

import (
	"testing"
)

func TestMatchToken(t *testing.T) {

	testString := "{}a[]"
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0}, // type must be not specified, it is inferred
		{RIGHT_CUR_BR, 1},
		{INVALID, 2},
		{LEFT_SQ_BR, 3},
		{RIGHT_SQ_BR, 4},
	}
	l := Lexer{testString, testString[0], 0}

	for l.Position < len(l.String) {
		token := l.matchToken()
		if token != expectedOutput[token.Position] {
			t.Errorf("\nThe input string: %v, position: %v \n"+
				"matches token %v which \n"+
				"does not match expected output %v\n",
				string(l.Character),
				l.Position,
				token,
				expectedOutput[l.Position])
		}
		l.nextElement()
	}
}
