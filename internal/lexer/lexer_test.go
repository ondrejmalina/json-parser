package lexer

import (
	"testing"
)

func TestMatchToken(t *testing.T) {

	testString := []rune("{}a[]")
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{RIGHT_CUR_BR, 1},
		{INVALID, 2},
		{LEFT_SQ_BR, 3},
		{RIGHT_SQ_BR, 4},
	}
	l := Lexer{testString, 0}

	for l.Position < len(l.Runes) {
		token := l.matchToken()
		if token != expectedOutput[token.Position] {
			t.Errorf("\nThe input string: %v, position: %v \n"+
				"matches token %v which \n"+
				"does not match expected output %v\n",
				string(l.Runes[l.Position]),
				l.Position,
				token,
				expectedOutput[l.Position])
		}
		l.nextElement()
	}
}

func TestTokenizeString(t *testing.T) {

	testString := []rune("{[1],}$")
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{LEFT_SQ_BR, 1},
		{DIGIT, 2},
		{RIGHT_SQ_BR, 3},
		{COMMA, 4},
		{RIGHT_CUR_BR, 5},
		{INVALID, 6},
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

	testString := []rune("")
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

func TestParseString(t *testing.T) {
	// TODO: Refactor the test
	testStringValid := []rune(`{"STR"}`)
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{STRING, 1},
		{RIGHT_CUR_BR, 6},
	}

	l := Lexer{testStringValid, 0}
	tokens := l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutput[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutput[token],
			)
		}
	}

	testStringInvalid := []rune(`{"STR`)
	expectedOutputInvalid := []Token{
		{LEFT_CUR_BR, 0},
		{INVALID, 1},
	}

	l = Lexer{testStringInvalid, 0}
	tokens = l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutputInvalid[token] {
			t.Errorf("\nToken %v does not match expected token in invalid string %v",
				tokens[token],
				expectedOutput[token],
			)
		}
	}

}

func TestInteger(t *testing.T) {
	testStringValid := []rune(`{123}`)
	expectedOutput := []Token{
		{LEFT_CUR_BR, 0},
		{DIGIT, 1},
		{RIGHT_CUR_BR, 4},
	}

	l := Lexer{testStringValid, 0}
	tokens := l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutput[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutput[token],
			)
		}
	}
}

func TestNull(t *testing.T) {
	testStringValid := []rune(`,null}`)
	testStringInvalid := []rune(`{nul`)
	expectedOutputValid := []Token{
		{COMMA, 0},
		{NULL, 1},
		{RIGHT_CUR_BR, 5},
	}
	expectedOutputInvalid := []Token{
		{LEFT_CUR_BR, 0},
		{INVALID, 1},
	}

	l := Lexer{testStringValid, 0}
	tokens := l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutputValid[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutputValid[token],
			)
		}
	}

	l = Lexer{testStringInvalid, 0}
	tokens = l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutputInvalid[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutputInvalid[token],
			)
		}
	}
}

func TestBool(t *testing.T) {
	testStringValid := []rune(`{false}`)
	testStringInvalid := []rune(`{fa`)
	expectedOutputValid := []Token{
		{LEFT_CUR_BR, 0},
		{BOOL, 1},
		{RIGHT_CUR_BR, 6},
	}
	expectedOutputInvalid := []Token{
		{LEFT_CUR_BR, 0},
		{INVALID, 1},
	}

	l := Lexer{testStringValid, 0}
	tokens := l.TokenizeString()
	t.Logf("%v", tokens)
	for token := range tokens {
		if tokens[token] != expectedOutputValid[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutputValid[token],
			)
		}
	}

	l = Lexer{testStringInvalid, 0}
	tokens = l.TokenizeString()
	for token := range tokens {
		if tokens[token] != expectedOutputInvalid[token] {
			t.Errorf("\nToken %v does not match expected token in valid string %v",
				tokens[token],
				expectedOutputInvalid[token],
			)
		}
	}
}
