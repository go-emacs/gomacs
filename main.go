package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/atotto/gomacs/internal/emacs"
)

var packages = emacs.List{
	emacs.Lisp("github.com/dominikh/go-mode.el"),
	emacs.Lisp("github.com/golang/lint/misc/emacs"),
	emacs.Lisp("github.com/syohex/emacs-go-eldoc"),
	emacs.Cmd("golang.org/x/tools/cmd/oracle"),
	emacs.Cmd("golang.org/x/tools/cmd/goimports"),
	emacs.Cmd("golang.org/x/tools/cmd/gorename"),
	emacs.Cmd("github.com/golang/lint/golint"),
	emacs.Cmd("github.com/nsf/gocode"),
	emacs.Cmd("code.google.com/p/rog-go/exp/cmd/godef"),
}

func main() {
	config := GetConfig()

	options := parse()

	exe, err := exec.LookPath(config.Emacs)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	env := os.Environ()
	args := emacsArgs(config.Args, options)
	err = syscall.Exec(exe, args, env)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
}

const gomacsOpt = "-gomacs."

func parse() (options []string) {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, gomacsOpt) {
			switch strings.TrimPrefix(arg, gomacsOpt) {
			case "update":
				packages.Update()
			case "init":
				packages.Install()
			default:
				panic("not implemented yet.")
			}
		} else {
			options = append(options, arg)
		}
	}
	return
}

func emacsArgs(config, options []string) []string {
	args := []string{"emacs"}
	args = append(args, config...)
	args = append(args, packages.Args()...)
	args = append(args, "-l", filepath.Join(EMACS_DIR, "env.el"))
	args = append(args, "-l", filepath.Join(EMACS_DIR, "init.el"))
	args = append(args, options...)

	return args
}
