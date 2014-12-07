package env

import (
	"go/build"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

const PACKAGE_PATH = "github.com/go-emacs/gomacs"

var GOMACS_DIR string
var ELISP_PATH string
var EMACSD_PATH string
var GOPATH string

func init() {
	p, err := build.Import(PACKAGE_PATH, build.Default.GOROOT, build.FindOnly)
	if err != nil {
		panic(err)
	}
	GOMACS_DIR = p.Dir
	EMACSD_PATH = filepath.Join(GOMACS_DIR, "emacs.d")
	ELISP_PATH = filepath.Join(EMACSD_PATH, "elisp")
	GOPATH = os.Getenv("GOPATH")
	if GOPATH == "" {
		log.Fatal("Please set GOPATH environment.")
	}
	GOPATH = filepath.SplitList(GOPATH)[0]

	generateEnvEL()
}

var (
	config map[string]string
	t      *template.Template
)

func generateEnvEL() {
	config = map[string]string{
		"GOROOT":             runtime.GOROOT(),
		"GOPATH":             GOPATH,
		"GOMACS_EMACSD_PATH": EMACSD_PATH,
	}
	t = template.Must(template.New("env.el").Parse(env_el_template))
	f, err := os.Create(filepath.Join(EMACSD_PATH, "env.el"))
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
