package util_test

import (
	"testing"

	"github.com/atotto/gomacs/util"
)

func TestGetConfig(t *testing.T) {
	c := util.GetConfig()

	if len(c.Args) == 0 || c.Args[0] != "-q" {
		t.Fatalf("want -q option got %+v", c.Args)
	}

	if c.Emacs != "emacs" {
		t.Fatalf("want emacs got %s", c.Emacs)
	}
}
