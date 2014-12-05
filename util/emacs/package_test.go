package emacs_test

import (
	"strings"
	"testing"

	"github.com/atotto/gomacs/util/emacs"
)

const testPackage = "github.com/atotto/yasnippet-golang"

func TestDownload(t *testing.T) {
	p := emacs.Lisp(testPackage)
	p.Install()
	p.Clean()
}

func TestPackagesArgs(t *testing.T) {
	p := emacs.Lisp(testPackage)
	p.Install()
	defer p.Clean()

	pkgs := emacs.List{emacs.Lisp(testPackage), emacs.Cmd("foo")}
	args := pkgs.Args()
	if args[0] != "-L" || !strings.HasSuffix(args[1], testPackage) {
		t.Fatalf("want -L and %s, got %+v", testPackage, args)
	}
}
