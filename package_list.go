package main

import (
	"github.com/atotto/gomacs/internal/emacs"
)

var packages = emacs.List{
	// for elisp
	emacs.Elisp("http://elpa.gnu.org/packages/cl-lib-0.5.el", "cl-lib.el"), // for emacs version <= 24.2

	// for go get
	emacs.ElispPackage("github.com/dominikh/go-mode.el"),
	emacs.ElispPackage("github.com/golang/lint/misc/emacs"),
	emacs.ElispPackage("github.com/syohex/emacs-go-eldoc"),
	emacs.ElispPackage("github.com/auto-complete/auto-complete"),
	emacs.ElispPackage("github.com/auto-complete/popup-el"),
	emacs.ElispPackage("github.com/capitaomorte/yasnippet"),
	emacs.ElispPackage("github.com/atotto/yasnippet-golang"),

	emacs.Cmd("golang.org/x/tools/cmd/oracle"),
	emacs.Cmd("golang.org/x/tools/cmd/goimports"),
	emacs.Cmd("golang.org/x/tools/cmd/gorename"),
	emacs.Cmd("github.com/golang/lint/golint"),
	emacs.Cmd("github.com/nsf/gocode"),
	emacs.Cmd("code.google.com/p/rog-go/exp/cmd/godef"),
}
