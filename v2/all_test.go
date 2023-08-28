// Copyright 2023 The Zlib-Go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z // import "modernc.org/z/v2"

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	util "modernc.org/ccgo/v3/lib"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	defer os.Remove("foo.gz")

	out, err := exec.Command("go", "run", filepath.Join("example", fmt.Sprintf("example_%s_%s.go", runtime.GOOS, runtime.GOARCH))).CombinedOutput()
	t.Logf("\n%s", out)
	if err != nil {
		t.Error(err)
	}
}

func Test2(t *testing.T) {
	tmpDir := t.TempDir()
	wd, err := util.AbsCwd()
	if err != nil {
		t.Fatal(err)
	}

	goos := runtime.GOOS
	goarch := runtime.GOARCH
	mg := filepath.Join(wd, "minigzip", fmt.Sprintf("minigzip_%s_%s.go", goos, goarch))
	ex := filepath.Join(wd, "example", fmt.Sprintf("example_%s_%s.go", goos, goarch))
	mgBin := "minigzip"
	exBin := "example"
	if goos == "windows" {
		mgBin += ".exe"
		exBin += ".exe"
	}
	if util.Shell("go", "build", "-o", filepath.Join(tmpDir, mgBin), mg); err != nil {
		t.Fatal(err)
	}

	if util.Shell("go", "build", "-o", filepath.Join(tmpDir, exBin), ex); err != nil {
		t.Fatal(err)
	}

	if err := util.InDir(tmpDir, func() error {
		switch goos {
		case "windows":
			if err := util.InDir(tmpDir, func() error {
				out, err := util.Shell("cmd.exe", "/c", fmt.Sprintf("echo hello world | %s | %[1]s -d", mgBin, exBin))
				if err != nil {
					return fmt.Errorf("%s\nFAIL: %v", out, err)
				}

				t.Logf("\n%s", out)
				return nil
			}); err != nil {
				t.Fatal(err)
			}
		default:
			if err := util.InDir(tmpDir, func() error {
				mgBin = "./" + mgBin
				exBin = "./" + exBin
				out, err := util.Shell("sh", "-c", fmt.Sprintf("echo hello world | %s | %[1]s -d && %s tmp", mgBin, exBin))
				if err != nil {
					return fmt.Errorf("%s\nFAIL: %v", out, err)
				}

				t.Logf("\n%s", out)
				return nil
			}); err != nil {
				t.Fatal(err)
			}
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
