package hsd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoSomething(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.dat")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range fns {
		t.Log(fn)

		b, err := os.ReadFile(fn)
		if err != nil {
			t.Fatal(err)
		}

		got := doSomething(string(b))

		b, err = os.ReadFile(fn[:len(fn)-4] + "out")
		if err != nil {
			t.Fatal(err)
		}
		want := string(b)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(diff)
		}
	}
}
