package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/ondrejmalina/json-parser/src/lexer"
)

func TestParser(t *testing.T) {

	cdir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get current working directory")
	}

	file, err := os.ReadFile(fmt.Sprintf("%v/testdata/step1/valid.json", cdir))
	if err != nil {
		t.Errorf("Failed to read the test file")
	}

	lex := lexer.Lexer{String: string(file), Character: file[0], Position: 0}
	tok := lex.TokenizeString()

	par := CreateParser(tok)
	err = par.parseValue()

	if err != nil {
		t.Errorf("Failed to parse valid json from step1")
	}
}
