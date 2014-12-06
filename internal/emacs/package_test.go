package emacs_test

import (
	"strings"
	"testing"

	"github.com/atotto/gomacs/internal/emacs"
)

const testPackage = "github.com/atotto/yasnippet-golang"

func TestDownload(t *testing.T) {
	p := emacs.ElispPackage(testPackage)
	p.Install()
	if p.IsInstaled() == false {
		t.Fatalf("want installed")
	}
	p.Clean()
	if p.IsInstaled() == true {
		t.Fatalf("want not installed")
	}
}

func TestPackagesArgs(t *testing.T) {
	p := emacs.ElispPackage(testPackage)
	p.Install()
	defer p.Clean()

	pkgs := emacs.List{emacs.ElispPackage(testPackage), emacs.Cmd("foo")}
	args := pkgs.Args()
	if args[0] != "-L" || !strings.HasSuffix(args[1], testPackage) {
		t.Fatalf("want -L and %s, got %+v", testPackage, args)
	}
}
