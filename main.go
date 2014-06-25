package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/atotto/gomacs/util"
	"github.com/atotto/gomacs/util/emacs"
)

var packages = emacs.Packages{
	emacs.Lisp("github.com/golang/lint/misc/emacs"),
	emacs.Lisp("github.com/syohex/emacs-go-eldoc"),
	emacs.Cmd("github.com/golang/lint/golint"),
}

func main() {
	config := util.GetConfig()

	options := parse()

	exe, err := exec.LookPath(config.Emacs)
	if err != nil {
		panic(err)
	}

	env := os.Environ()
	args := emacsArgs(config.Args, options)
	err = syscall.Exec(exe, args, env)
	if err != nil {
		panic(err)
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
				packages.Setup()
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
	args = append(args, util.GoModeLoadPath()...)
	args = append(args, "-l", filepath.Join(util.Emacsd(), "define.el"))
	args = append(args, "-l", filepath.Join(util.Emacsd(), "init.el"))
	args = append(args, options...)

	return args
}
