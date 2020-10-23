package vfscopy

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"testing"
	"errors"

	"git.rootprojects.org/root/vfscopy/fixtures"
)

func TestNativeRecursiveCopy(t *testing.T) {
	{
		opts := Options{
			OnSymlink: func(path string) SymlinkAction {
				return Shallow
			},
		}
		vfs := Dir("./fixtures/fs/")
		tmpDir := "/tmp/go-vfscopy-dst-1/"

		_ = os.RemoveAll(tmpDir)
		defer func () {
			_ = os.RemoveAll(tmpDir)
		}()

		if err := CopyAll(vfs, ".", tmpDir, opts); nil != err {
			t.Errorf("error: %v", err)
			return
		}
	}
}

func TestVFSRecursiveCopy(t *testing.T) {
	{
		opts := Options{
			OnSymlink: func(path string) SymlinkAction {
				return Shallow
			},
		}
		vfs := NewVFS(fixtures.Assets)
		tmpDir := "/tmp/go-vfscopy-dst-2/"

		_ = os.RemoveAll(tmpDir)
		defer func () {
			_ = os.RemoveAll(tmpDir)
		}()

		if err := CopyAll(vfs, ".", tmpDir, opts); nil != err {
			t.Errorf("copy error: %v", err)
			return
		}
		root := "fixtures/fs"
		walker := func(path string, info os.FileInfo, err error) error {
			if nil != err {
				return err
			}

			rel := path[len(root):]

			s, err := os.Lstat(path)
			if nil != err {
				return err
			}

			dst := filepath.Join(tmpDir, rel)
			d, err := os.Lstat(dst)
			if nil != err {
				return err
			}

			if s.Mode() != d.Mode() {
				log.Println("implementation does not support permissions and/or symlinks:", s.Mode(), d.Mode())
				//return errors.New("did not copy file mode (e.g. symlink status)")
			} else if s.Size() != d.Size() {
				return errors.New("did not copy full file size")
			}

			fmt.Println(path, "=>", dst)
			return nil
		}
		err := filepath.Walk(root, walker)
		if nil != err {
			t.Errorf("check dst error: %v", err)
			return
		}
	}
}
