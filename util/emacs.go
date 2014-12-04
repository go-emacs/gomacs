package util

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/atotto/gomacs/util/emacs"
)

var gomacsDir string

func init() {
	gomacsDir = emacs.Cmd("github.com/atotto/gomacs").LocalPath()
	genDefineel()
}

//
func GomacsDir() string {
	return gomacsDir
}

//
func ElispDir() string {
	return filepath.Join(gomacsDir, "elisp")
}

//
func Emacsd() string {
	return filepath.Join(gomacsDir, "emacs.d")
}

//
func GetLoadPaths() []string {
	return nil
}

var (
	config map[string]string
	t      *template.Template
)

func genDefineel() {
	config = map[string]string{
		"GOROOT":             runtime.GOROOT(),
		"GOPATH":             os.Getenv("GOPATH"), // TODO: parse first path
		"GOMACS_EMACSD_PATH": Emacsd(),
	}
	t = template.Must(template.New("define.el").Parse(initel))
	f, err := os.Create(filepath.Join(Emacsd(), "define.el"))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, config)
	if err != nil {
		log.Fatal(err)
	}
}

const initel = `
(defvar goroot "{{.GOROOT}}")
(defvar gopath "{{.GOPATH}}")
(defvar gomacs-emacsd-path "{{.GOMACS_EMACSD_PATH}}")
`
