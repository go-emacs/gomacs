package util

import (
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/atotto/gomacs/util/emacs"
)

var gomacsDir string

func init() {
	gomacsDir = emacs.Cmd("github.com/atotto/gomacs").LocalPath()
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

//
func GoModeLoadPath() []string {
	return []string{"-L", filepath.Join(runtime.GOROOT(), "misc", "emacs")}
}

func genInitel() {
	config = map[string]string{
		"GOROOT": runtime.GOROOT(),
	}
	t = template.Must(template.New("init.el").Parse(initel))
}

var (
	config map[string]string
	t      *template.Template
)

const initel = `

`
