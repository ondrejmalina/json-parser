package parser

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ondrejmalina/json-parser/internal/cli"
	"github.com/ondrejmalina/json-parser/internal/lexer"
)

func TestParsing(t *testing.T) {

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
		{fmt.Sprintf("%v/testdata/step1/invalid2.json", cdir), false},
		{fmt.Sprintf("%v/testdata/step2/valid.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step2/valid2.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step2/valid3.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step2/invalid.json", cdir), false},
		{fmt.Sprintf("%v/testdata/step2/invalid2.json", cdir), false},
		{fmt.Sprintf("%v/testdata/step3/valid.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step3/invalid.json", cdir), false},
		{fmt.Sprintf("%v/testdata/step3/invalid2.json", cdir), false},
		{fmt.Sprintf("%v/testdata/step4/valid.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step4/valid2.json", cdir), true},
		{fmt.Sprintf("%v/testdata/step4/invalid.json", cdir), false},
	}

	for _, testFile := range testFiles {

		// input
		file, err := os.ReadFile(testFile.testFile)
		if err != nil {
			t.Errorf("\nFailed to read the test file")
		}
		userInput := cli.UserInput{Input: string(file)}

		// lexer
		lexer := lexer.CreateLexer(userInput)

		// parser
		p := CreateParser(*lexer)
		err = p.ParseJson()

		// parsing validation
		switch {
		case err == nil && testFile.valid == false:
			t.Errorf("Current token %v", p.token)
			t.Errorf("Invalid file %v parsed as valid, \nerror: %v", testFile.testFile, err)
		case err != nil && testFile.valid == true:
			t.Errorf("Current token %v", p.token)
			t.Errorf("Valid file %v parsed as invalid, \nerror: %v", testFile.testFile, err)
		case err == nil && testFile.valid == true:
			fileParent := strings.Split(testFile.testFile, "/")
			t.Logf("File %v/%v parsed successfully", fileParent[len(fileParent)-2], fileParent[len(fileParent)-1])
		case err != nil && testFile.valid == false:
			fileParent := strings.Split(testFile.testFile, "/")
			t.Logf("File %v/%v parsed successfully", fileParent[len(fileParent)-2], fileParent[len(fileParent)-1])
		}
	}
}
