package emacs

import (
	"bytes"
	"errors"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"
)

type kind int

const (
	cmd kind = 1 << iota
	lisp
)

// P represents package info.
type P struct {
	path string
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

func Lisp(path string) *P {
	return &P{path: path, kind: lisp}
}

func (p *P) Args() []string {
	if p.kind == lisp {
		return []string{"-L", p.LocalPath()}
	}
	return nil
}

// Install installs the package from internet.
func (p *P) Install() error {
	return fetch(p.path, false)
}

func (p *P) IsInstaled() bool {
	if p.kind == cmd {
		_, err := exec.LookPath(path.Base(p.path))
		return err == nil
	} else {
		_, err := build.Import(p.path, os.Getenv("GOPATH"), build.FindOnly)
		return err == nil
	}
}

// Update updates the package from internet.
func (p *P) Update() error {
	// TODO: implement
	return fetch(p.path, true)
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

func fetch(pkgpath string, update bool) error {
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
