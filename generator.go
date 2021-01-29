// Copyright 2021 The Zlib-Go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	tarFn  = tarVer + ".tar.gz"
	tarVer = "zlib-1.2.11"
)

type supportedKey = struct{ os, arch string }

var (
	goos      = env("TARGET_GOOS", runtime.GOOS)
	goarch    = env("TARGET_GOARCH", runtime.GOARCH)
	supported = map[supportedKey]struct{}{
		{"darwin", "amd64"}: {},
		{"linux", "386"}:    {},
		{"linux", "amd64"}:  {},
		{"linux", "arm"}:    {},
		{"linux", "arm64"}:  {},
		//TODO {"windows", "386"}:   {},
		//TODO {"windows", "amd64"}: {},
	}
	tmpDir = os.Getenv("GO_GENERATE_TMPDIR")
	ccgo   string
)

func env(key, dflt string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}

	return dflt
}

func fatalf(s string, args ...interface{}) {
	s = strings.TrimRight(fmt.Sprintf(s, args...), " \t\n\r")
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func fatal(args ...interface{}) {
	s := strings.TrimRight(fmt.Sprint(args...), " \t\n\r")
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	if _, ok := supported[supportedKey{goos, goarch}]; !ok {
		fatalf("unknown/unsupported os/arch combination: %s/%s\n", goos, goarch)
	}

	var err error
	if ccgo, err = exec.LookPath("ccgo"); err != nil {
		fatal(err)
	}

	if err := mkdir("lib"); err != nil {
		fatal(err)
	}

	if err := mkdir("internal"); err != nil {
		fatal(err)
	}

	if tmpDir == "" {
		var err error
		if tmpDir, err = ioutil.TempDir("", "go-generate-zlib-"); err != nil {
			fatal(err)
		}

		defer os.RemoveAll(tmpDir)
	}

	f, err := os.Open(tarFn)
	if err != nil {
		fatal(err)
	}

	defer f.Close()

	if err := inDir(tmpDir, func() error {
		os.RemoveAll(tarVer)
		if err := untar("", bufio.NewReader(f)); err != nil {
			fatal(err)
		}

		return nil
	}); err != nil {
		fatal(err)
	}

	cdb := filepath.Join(tmpDir, "cdb.json")
	if err := inDir(filepath.Join(tmpDir, tarVer), func() error {
		if _, err := shell("./configure", "--static", "--64"); err != nil {
			fatal(err)
		}

		if _, err := shell("ccgo", "-compiledb", cdb, "make", "test64"); err != nil {
			fatal(err)
		}

		return nil
	}); err != nil {
		fatal(err)
	}
	mustCC(os.Stdout, os.Stderr,
		"-export-defines", "",
		"-export-enums", "",
		"-export-externs", "X",
		"-export-fields", "F",
		"-export-structs", "",
		"-export-typedefs", "",
		"-o", filepath.Join("lib", fmt.Sprintf("z_%s_%s.go", goos, goarch)),
		"-pkgname", "z",
		"-trace-translation-units",
		cdb, "libz.a",
	)
	mustCC(os.Stdout, os.Stderr,
		"-lmodernc.org/z/lib",
		"-o", filepath.Join("internal", fmt.Sprintf("minigzip_%s_%s.go", goos, goarch)),
		"-trace-translation-units",
		cdb, "minigzip64",
	)
	mustCC(os.Stdout, os.Stderr,
		"-lmodernc.org/z/lib",
		"-o", filepath.Join("internal", fmt.Sprintf("example_%s_%s.go", goos, goarch)),
		"-trace-translation-units",
		cdb, "example64",
	)
}

func mkdir(paths ...string) error {
	for _, path := range paths {
		path = filepath.FromSlash(path)
		if err := os.MkdirAll(path, 0770); err != nil {
			return err
		}
	}
	return nil
}

func inDir(dir string, f func() error) (err error) {
	var cwd string
	if cwd, err = os.Getwd(); err != nil {
		return err
	}

	defer func() {
		if err2 := os.Chdir(cwd); err2 != nil {
			err = err2
		}
	}()

	if err = os.Chdir(dir); err != nil {
		return err
	}

	return f()
}

func untar(root string, r io.Reader) error {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err != io.EOF {
				return err
			}

			return nil
		}

		switch hdr.Typeflag {
		case tar.TypeDir:
			dir := filepath.Join(root, hdr.Name)
			if err = os.MkdirAll(dir, 0770); err != nil {
				return err
			}
		case tar.TypeSymlink, tar.TypeXGlobalHeader:
			// skip
		case tar.TypeReg, tar.TypeRegA:
			dir := filepath.Dir(filepath.Join(root, hdr.Name))
			if _, err := os.Stat(dir); err != nil {
				if !os.IsNotExist(err) {
					return err
				}

				if err = os.MkdirAll(dir, 0770); err != nil {
					return err
				}
			}

			f, err := os.OpenFile(filepath.Join(root, hdr.Name), os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			w := bufio.NewWriter(f)
			if _, err = io.Copy(w, tr); err != nil {
				return err
			}

			if err := w.Flush(); err != nil {
				return err
			}

			if err := f.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected tar header typeflag %#02x", hdr.Typeflag)
		}
	}
}

type echoWriter struct {
	w bytes.Buffer
}

func (w *echoWriter) Write(b []byte) (int, error) {
	os.Stdout.Write(b)
	return w.w.Write(b)
}

func shell(cmd string, args ...string) ([]byte, error) {
	wd, err := absCwd()
	if err != nil {
		return nil, err
	}

	fmt.Printf("execute %s %q in %s\n", cmd, args, wd)
	var b echoWriter
	c := exec.Command(cmd, args...)
	c.Stdout = &b
	c.Stderr = &b
	err = c.Run()
	return b.w.Bytes(), err
}

func absCwd() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if wd, err = filepath.Abs(wd); err != nil {
		return "", err
	}

	return wd, nil
}

func cc(stdout, stderr io.Writer, args ...string) error {
	cmd := exec.Command(ccgo, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}

func mustCC(stdout, stderr io.Writer, args ...string) {
	if err := cc(stdout, stderr, args...); err != nil {
		fatal(err)
	}
}
