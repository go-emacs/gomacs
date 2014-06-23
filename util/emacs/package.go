package emacs

import (
	"bytes"
	"errors"
	"go/build"
	"log"
	"os"
	"os/exec"
)

type kind int

const (
	cmd kind = 1 << iota
	lisp
)

type Package struct {
	path string
	kind kind
}

type Packages []*Package

func (pkgs *Packages) iter(fn func(*Package) error) error {
	for _, pkg := range *pkgs {
		log.Printf("fetch: %s\n", pkg.path)
		err := fn(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pkgs *Packages) Setup() error {
	return pkgs.iter((*Package).Setup)
}

func (pkgs *Packages) Update() error {
	return pkgs.iter((*Package).Update)
}

func (pkgs *Packages) Args() (args []string) {
	for _, pkg := range *pkgs {
		arg := pkg.Args()
		if arg != nil {
			args = append(args, arg...)
		}
	}
	return args
}

func Cmd(path string) *Package {
	return &Package{path: path, kind: cmd}
}

func Lisp(path string) *Package {
	return &Package{path: path, kind: lisp}
}

func (pkg *Package) Args() []string {
	if pkg.kind == lisp {
		return []string{"-L", pkg.LocalPath()}
	}
	return nil
}

// Download downloads the package from internet.
func (pkg *Package) Setup() error {
	return fetch(pkg.path, false)
}

// Update updates the package from internet.
func (pkg *Package) Update() error {
	// TODO: implement
	return fetch(pkg.path, true)
}

// Clean remove the package from file system.
func (pkg *Package) Clean() error {
	return clean(pkg.path)
}

// LocalPath returns the package local absolute path.
func (pkg *Package) LocalPath() string {
	path, err := packageAbsolutePath(pkg.path)
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
