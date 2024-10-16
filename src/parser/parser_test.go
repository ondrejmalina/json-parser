package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/ondrejmalina/json-parser/src/lexer"
)

func TestParser(t *testing.T) {

	type testFile struct {
		testFile string
		valid    bool
	}

	cdir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get current working directory")
	}

	testFiles := []testFile{
		{fmt.Sprintf("%v/testdata/step1/valid.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step1/invalid.json", cdir), false},
	}

	for _, testFile := range testFiles {

		file, err := os.ReadFile(testFile.testFile)
		if err != nil {
			t.Errorf("Failed to read the test file")
		}

		lex := lexer.Lexer{String: string(file), Position: 0}
		tok := lex.TokenizeString()

		par := CreateParser(tok)
		err = par.parseValue()

		switch err {
		case nil:
			if testFile.valid == false {
				t.Errorf("Invalid file %v parsed as valid", testFile.testFile)
			}
		default:
			if testFile.valid == true {
				t.Errorf("Valid file %v parsed as invalid", testFile.testFile)
			}
		}
	}
}
