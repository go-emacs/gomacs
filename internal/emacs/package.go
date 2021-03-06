package emacs

import (
	"bytes"
	"errors"
	"go/build"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-emacs/gomacs/internal/env"
)

type kind int

const (
	cmd kind = 1 << iota
	lisp
	lispfile
)

type PackageManager interface {
	Install() error
	InstallForce() error
	Update() error
	Clean() error
}

// P represents package info.
type P struct {
	path string
	name string
	kind kind
}

type List []*P

func (l *List) iter(fn func(*P) error) error {
	for _, pkg := range *l {
		log.Printf("install: %s\n", pkg.path)
		err := fn(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *List) InstallForce() error {
	return l.iter((*P).Install)
}

func (l *List) Install() error {
	for _, pkg := range *l {
		if pkg.IsInstaled() {
			continue
		}
		log.Printf("install: %s\n", pkg.path)
		err := pkg.Install()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *List) Update() error {
	return l.iter((*P).Update)
}

func (l *List) Args() (args []string) {
	for _, pkg := range *l {
		arg := pkg.Args()
		if arg != nil {
			args = append(args, arg...)
		}
	}
	return args
}

func Cmd(path string) *P {
	return &P{path: path, kind: cmd}
}

func ElispPackage(path string) *P {
	if strings.HasPrefix(path, "http://") {
		return &P{path: path, kind: lispfile}
	}
	return &P{path: path, kind: lisp}
}

func Elisp(url string, name string) *P {
	if strings.HasPrefix(url, "http://") {
		if name == "" {
			name = path.Base(url)
		}
		return &P{path: url, name: name, kind: lispfile}
	}
	return &P{path: url, kind: lisp}
}

func (p *P) Args() []string {
	if p.kind == lisp {
		return []string{"-L", p.LocalPath()}
	}
	return nil
}

// Install installs the package from internet.
func (p *P) Install() error {
	switch p.kind {
	case lispfile:
		return wget(p.path, p.name)
	default:
		return goget(p.path, false)
	}
}

func (p *P) IsInstaled() bool {
	switch p.kind {
	case cmd:
		_, err := exec.LookPath(path.Base(p.path))
		return err == nil
	case lisp:
		_, err := build.Import(p.path, build.Default.GOROOT, build.FindOnly)
		return err == nil
	case lispfile:
		_, err := os.Stat(filepath.Join(env.ELISP_PATH, p.name))
		return !os.IsNotExist(err)
	default:
		panic("not implemeted yet.")
	}
}

// Update updates the package from internet.
func (p *P) Update() error {
	// TODO: implement
	return goget(p.path, true)
}

// Clean remove the package from file system.
func (p *P) Clean() error {
	return clean(p.path)
}

// LocalPath returns the package local absolute path.
func (p *P) LocalPath() string {
	path, err := packageAbsolutePath(p.path)
	if err != nil {
		log.Fatal(err)
	}
	return path
}

var no_error = []byte(`no buildable Go source files`)

func goget(pkgpath string, update bool) error {
	var args []string
	if update {
		args = []string{"get", "-u", pkgpath}
	} else {
		args = []string{"get", pkgpath}
	}
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout

	var buf bytes.Buffer
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		if !bytes.Contains(buf.Bytes(), no_error) {
			return errors.New(buf.String())
		}
	}
	return nil
}

func wget(url string, name string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = resp.Body.Close(); err != nil {
		return
	}
	fname := filepath.Join(env.ELISP_PATH, name)
	// truncate file if it already exists.
	if err = ioutil.WriteFile(fname, buf, 0666); err != nil {
		return
	}
	return nil
}

func clean(pkgpath string) error {
	p, err := build.Import(pkgpath, build.Default.GOROOT, build.FindOnly)
	if err != nil {
		return err
	}
	os.RemoveAll(p.Dir)
	return nil
}

func packageAbsolutePath(pkg string) (string, error) {
	p, err := build.Import(pkg, build.Default.GOROOT, build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, nil
}
