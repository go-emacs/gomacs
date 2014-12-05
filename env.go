package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/atotto/gomacs/internal/emacs"
)

var GOMACS_DIR string
var ELISP_DIR string
var EMACS_DIR string

func init() {
	GOMACS_DIR = emacs.Cmd("github.com/atotto/gomacs").LocalPath()
	ELISP_DIR = filepath.Join(GOMACS_DIR, "elisp")
	EMACS_DIR = filepath.Join(GOMACS_DIR, "emacs.d")

	generateEnvEL()
}

var (
	config map[string]string
	t      *template.Template
)

func generateEnvEL() {
	config = map[string]string{
		"GOROOT":             runtime.GOROOT(),
		"GOPATH":             os.Getenv("GOPATH"), // TODO: parse first path
		"GOMACS_EMACSD_PATH": EMACS_DIR,
	}
	t = template.Must(template.New("env.el").Parse(env_el_template))
	f, err := os.Create(filepath.Join(EMACS_DIR, "env.el"))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, config)
	if err != nil {
		log.Fatal(err)
	}
}

const env_el_template = `
(defvar goroot "{{.GOROOT}}")
(defvar gopath "{{.GOPATH}}")
(defvar gomacs-emacsd-path "{{.GOMACS_EMACSD_PATH}}")
`
