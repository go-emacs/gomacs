package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	options := Parse()

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

func Parse() (options []string) {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--help":
			Usage()
			options = append(options, arg)
			break
		case "--update":
			packages.Update()
			os.Exit(0)
		case "--install":
			packages.InstallForce()
			os.Exit(0)
		default:
			options = append(options, arg)
		}
	}
	packages.Install()
	return
}

var usage = `
Launch:

   $ gomacs                  # launch emacs

Update:

   $ gomacs --update         # update emacs lisp from internet.

The gomacs can use emacs option and operation. for example:

   $ gomacs --help           # show emacs --help
   $ gomacs main.go          # open main.go
   $ gomacs +12 main.go      # go to line
   $ gomacs -rv              # switch foreground and background color

`

func Usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Println(usage)
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
