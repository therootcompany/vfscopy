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
		if err := Copy( "./fixtures/src/", "/tmp/dst/", opts,); nil != err {
			t.Fatalf("error: %v", err)
		}
	}
}
