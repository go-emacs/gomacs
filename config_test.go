package main

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	c := GetConfig()

	if len(c.Args) == 0 || c.Args[0] != "-q" {
		t.Fatalf("want -q option got %+v", c.Args)
	}

	if c.Emacs != "emacs" {
		t.Fatalf("want emacs got %s", c.Emacs)
	}
}
