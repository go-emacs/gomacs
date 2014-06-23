package util_test

import (
	"strings"
	"testing"

	"github.com/atotto/gomacs/util"
)

func TestElispDir(t *testing.T) {
	actual := util.ElispDir()

	expected := "github.com/atotto/gomacs/elisp"

	if !strings.Contains(actual, expected) {
		t.Fatalf("want containing %s, got %s", expected, actual)
	}
}
