package strings_test

import (
	"testing"

	"github.com/Kamaalio/kamaalgo/strings"
)

func TestUnwrapEnvironment(t *testing.T) {
	value, err := strings.Unwrap("secret")
	if err != nil {
		t.Errorf("Result was incorrect, got: %v, want: %v.", err, nil)
	}
	if value != "secret" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", value, "secret")
	}
}

func TestUnwrapEnvironmentNotFound(t *testing.T) {
	_, err := strings.Unwrap("")
	if err != strings.ErrStringUnwrap {
		t.Errorf("Result was incorrect, got: %v, want: %v.", err, strings.ErrStringUnwrap)
	}
}
