package vfscopy

import (
	"testing"
)

func TestRecursiveCopy(t *testing.T) {
	{
		opts := Options{
			OnSymlink: func (path string) SymlinkAction {
				return Shallow
			},
		}
		vfs := Dir("./fixtures/src/")
		if err := Copy(vfs, ".", "/tmp/dst/", opts,); nil != err {
			t.Fatalf("error: %v", err)
		}
	}
}
