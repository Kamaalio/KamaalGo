package strings_test

import (
	"testing"

	"github.com/Kamaalio/kamaalgo/strings"
)

type testCase struct {
	Input          string
	ExpectedOutput string
}

func TestPascalToSnakeCase(t *testing.T) {
	testCases := []testCase{
		{Input: "TargetLocale", ExpectedOutput: "target_locale"},
		{Input: "Target", ExpectedOutput: "target"},
	}

	for _, singleTestCase := range testCases {
		result := strings.PascalToSnakeCase(singleTestCase.Input)
		if strings.PascalToSnakeCase(singleTestCase.Input) != singleTestCase.ExpectedOutput {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, singleTestCase.ExpectedOutput)
		}
	}
}
