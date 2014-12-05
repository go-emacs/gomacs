package main

import (
	"strings"
	"testing"
)

func TestElispDir(t *testing.T) {
	expected := "github.com/atotto/gomacs/elisp"

	if !strings.Contains(ELISP_DIR, expected) {
		t.Fatalf("want containing %s, got %s", expected, ELISP_DIR)
	}
}
