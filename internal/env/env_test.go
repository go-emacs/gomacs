package env

import (
	"strings"
	"testing"
)

func TestElispDir(t *testing.T) {
	expected := PACKAGE_PATH + "/emacs.d/elisp"

	if !strings.Contains(ELISP_PATH, expected) {
		t.Fatalf("want containing %s, got %s", expected, ELISP_PATH)
	}
}
