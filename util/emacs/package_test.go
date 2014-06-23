package emacs_test

import (
	"strings"
	"testing"

	"github.com/atotto/gomacs/util/emacs"
)

const testPackage = "github.com/atotto/yasnippet-golang"

func TestDownload(t *testing.T) {
	p := emacs.Lisp(testPackage)
	p.Setup()
	p.Clean()
}

func TestPackagesArgs(t *testing.T) {
	p := emacs.Lisp(testPackage)
	p.Setup()
	defer p.Clean()

	pkgs := emacs.Packages{emacs.Lisp(testPackage), emacs.Cmd("foo")}
	args := pkgs.Args()
	if args[0] != "-L" || !strings.HasSuffix(args[1], testPackage) {
		t.Fatalf("want -L and %s, got %+v", testPackage, args)
	}
}
