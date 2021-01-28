// Copyright 2021 The Zlib-Go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z // import "modernc.org/z"

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"modernc.org/ccgo/v3/lib"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func stack() string { return string(debug.Stack()) }

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO, stack) //TODOOK
}

// ----------------------------------------------------------------------------

const (
	tarFn  = tarVer + ".tar.gz"
	tarVer = "zlib-1.2.11"
)

type supportedKey = struct{ os, arch string }

var (
	oGenerate = flag.Bool("generate", false, "")

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
)

func env(key, dflt string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}

	return dflt
}

func TestMain(m *testing.M) {
	flag.Parse()
	rc := m.Run()
	os.Exit(rc)
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
	return ccgo.NewTask(append([]string{"ccgo"}, args...), stdout, stderr).Main()
}

func mustCC(t *testing.T, stdout, stderr io.Writer, args ...string) {
	if err := cc(stdout, stderr, args...); err != nil {
		t.Fatal(err)
	}
}

func TestGenerate(t *testing.T) {
	if !*oGenerate {
		t.Skip("not enabled")
	}

	if _, ok := supported[supportedKey{goos, goarch}]; !ok {
		t.Fatalf("unknown/unsupported os/arch combination: %s/%s\n", goos, goarch)
	}

	if err := mkdir("lib"); err != nil {
		t.Fatal(err)
	}

	if err := mkdir("internal"); err != nil {
		t.Fatal(err)
	}

	if tmpDir == "" {
		var err error
		if tmpDir, err = ioutil.TempDir("", "go-generate-zlib-"); err != nil {
			t.Fatal(err)
		}

		fmt.Println(tmpDir)
		//defer os.RemoveAll(tmpDir)
	}

	f, err := os.Open(tarFn)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	if err := inDir(tmpDir, func() error {
		os.RemoveAll(tarVer)
		if err := untar("", bufio.NewReader(f)); err != nil {
			t.Fatal(err)
		}

		if err := os.Chdir(tarVer); err != nil {
			t.Fatal(err)
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}

	cdb := filepath.Join(tmpDir, "cdb.json")
	switch {
	case goos == "darwin" && goarch == "amd64":
		b, err := ioutil.ReadFile("lib/posix.json")
		if err != nil {
			t.Fatal(err)
		}

		b = bytes.ReplaceAll(b, []byte("$TMP"), []byte(tmpDir))
		if err := ioutil.WriteFile(cdb, b, 0660); err != nil {
			t.Fatal(err)
		}
	default:
		if err := inDir(tmpDir, func() error {
			if _, err := shell("./configure", "--static", "--64"); err != nil {
				t.Fatal(err)
			}

			if _, err := shell("ccgo", "-compiledb", cdb, "make", "test64"); err != nil {
				t.Fatal(err)
			}

			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	mustCC(t, os.Stdout, os.Stderr,
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
	mustCC(t, os.Stdout, os.Stderr,
		"-lmodernc.org/z/lib",
		"-o", filepath.Join("internal", fmt.Sprintf("minigzip_%s_%s.go", goos, goarch)),
		"-trace-translation-units",
		cdb, "minigzip64",
	)
	mustCC(t, os.Stdout, os.Stderr,
		"-lmodernc.org/z/lib",
		"-o", filepath.Join("internal", fmt.Sprintf("example_%s_%s.go", goos, goarch)),
		"-trace-translation-units",
		cdb, "example64",
	)
}

func Test(t *testing.T) {
	if *oGenerate {
		return
	}

	tmpDir, err := ioutil.TempDir("", "go-test-zlib-")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tmpDir)

	wd, err := absCwd()
	if err != nil {
		t.Fatal(err)
	}

	goos := runtime.GOOS
	goarch := runtime.GOARCH
	mg := filepath.Join(wd, "internal", fmt.Sprintf("minigzip_%s_%s.go", goos, goarch))
	ex := filepath.Join(wd, "internal", fmt.Sprintf("example_%s_%s.go", goos, goarch))
	if err := inDir(tmpDir, func() error {
		if out, err := shell("sh", "-c", fmt.Sprintf("echo hello world | go run %s | go run %[1]s -d && go run %s tmp", mg, ex)); err != nil {
			return fmt.Errorf("%s\nFAIL: %v", out, err)
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
