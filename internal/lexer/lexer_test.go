package lexer

import (
	"testing"
)

func TestMatchStringToken(t *testing.T) {

	testRunes := []rune("\"ab\"")
	l := Lexer{runes: testRunes}

	tokString := Token{STRING, 0}
	if tok := l.GetToken(); tok != tokString {
		t.Errorf("String not parsed")
	}
}
