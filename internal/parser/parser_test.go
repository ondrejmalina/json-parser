package parser

//
// import (
// 	"fmt"
// 	"os"
// 	"testing"
//
// 	"github.com/ondrejmalina/json-parser/internal/lexer"
// )
//
// func TestArrayParsing(t *testing.T) {
// 	test_cases := []string{`[]`, `["Hello"]`, `[123]`, `[true]`, `[null]`}
//
// 	for _, c := range test_cases {
// 		lex := lexer.Lexer{Runes: []rune(c), CurrentPosition: 0}
// 		tok := lex.TokenizeString()
// 		t.Log(tok)
// 		p := CreateParser(tok)
// 		if err := p.parseArray(); err != nil {
// 			t.Errorf("Parsing of array failed")
// 		}
// 	}
// }
//
// func TestParser(t *testing.T) {
//
// 	type testFile struct {
// 		testFile string
// 		valid    bool
// 	}
//
// 	cdir, err := os.Getwd()
// 	if err != nil {
// 		t.Errorf("Failed to get current working directory")
// 	}
//
// 	testFiles := []testFile{
// 		{fmt.Sprintf("%v/testdata/step1/valid.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step1/invalid.json", cdir), false},
// 		{fmt.Sprintf("%v/testdata/step2/valid.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step2/valid2.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step2/valid3.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step2/invalid.json", cdir), false},
// 		{fmt.Sprintf("%v/testdata/step2/invalid2.json", cdir), false},
// 		{fmt.Sprintf("%v/testdata/step3/valid.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step3/invalid.json", cdir), false},
// 		{fmt.Sprintf("%v/testdata/step4/valid.json", cdir), true},
// 		{fmt.Sprintf("%v/testdata/step4/valid2.json", cdir), true},
// 		// {fmt.Sprintf("%v/testdata/step4/invalid.json", cdir), true},
// 	}
//
// 	for _, testFile := range testFiles {
//
// 		file, err := os.ReadFile(testFile.testFile)
// 		if err != nil {
// 			t.Errorf("\nFailed to read the test file")
// 		}
//
// 		lex := lexer.Lexer{Runes: []rune(string(file)), CurrentPosition: 0}
// 		tok := lex.TokenizeString()
//
// 		par := CreateParser(tok)
// 		err = par.parseValue()
//
// 		switch err {
// 		case nil:
// 			if testFile.valid == false {
// 				t.Errorf("Invalid file %v parsed as valid, \nerror: %v", testFile.testFile, err)
// 			}
// 		default:
// 			if testFile.valid == true {
// 				t.Errorf("Valid file %v parsed as invalid, \nerror: %v", testFile.testFile, err)
// 			}
// 		}
// 	}
// }
