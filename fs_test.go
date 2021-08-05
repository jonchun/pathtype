package pathtype_test

import (
	"io/fs"
	"testing"
)

func TestValidPath(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(fs.ValidPath(string(p)))
		t1.Result(p.ValidPath())
		t1.AssertEquals()
	}
}

func TestReadDirFS(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(fs.ReadDir(fsys, string(p)))
		t1.Result(p.ReadDirFS(fsys))
		t1.AssertSimilar()
	}
}

func TestSub(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(fs.Sub(fsys, string(p)))
		t1.Result(p.Sub(fsys))
		t1.AssertSimilar()
	}
}
