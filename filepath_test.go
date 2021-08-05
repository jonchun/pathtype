package pathtype_test

import (
	"path/filepath"
	"testing"
)

func TestAbs(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Abs(string(p)))
		t1.Result(p.Abs())
		t1.AssertEquals()
	}
}

func TestBase(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Base(string(p)))
		t1.Result(p.Base())
		t1.AssertEquals()
	}
}

func TestClean(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Clean(string(p)))
		t1.Result(p.Clean())
		t1.AssertEquals()
	}
}

func TestDir(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Dir(string(p)))
		t1.Result(p.Dir())
		t1.AssertEquals()
	}
}

func TestEvalSymlinks(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.EvalSymlinks(string(p)))
		t1.Result(p.EvalSymlinks())
		t1.AssertEquals()
	}
}

func TestExt(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Ext(string(p)))
		t1.Result(p.Ext())
		t1.AssertEquals()
	}
}

func TestFromSlash(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.FromSlash(string(p)))
		t1.Result(p.FromSlash())
		t1.AssertEquals()
	}
}

func TestGlob(t *testing.T) {
	for _, p := range testPaths {
		for _, p1 := range testPatterns {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(filepath.Glob(filepath.Join(string(p), p1)))
			t1.Result(p.Glob(p1))
			t1.AssertEquals()
		}
	}
}

func TestIsAbs(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.IsAbs(string(p)))
		t1.Result(p.IsAbs())
		t1.AssertEquals()
	}
}

func TestMatch(t *testing.T) {
	for _, p := range testPaths {
		for _, p1 := range testPatterns {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(filepath.Match(p1, string(p)))
			t1.Result(p.Match(p1))
			t1.AssertEquals()
		}
	}
}

func TestRel(t *testing.T) {
	for _, p := range testPaths {
		for _, p1 := range testPaths {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(filepath.Rel(string(p), string(p1)))
			t1.Result(p.Rel(p1))
			t1.AssertEquals()
		}
	}
}

func TestSplit(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.Split(string(p)))
		t1.Result(p.Split())
		t1.AssertEquals()
	}
}

func TestToSlash(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.ToSlash(string(p)))
		t1.Result(p.ToSlash())
		t1.AssertEquals()
	}
}
func TestVolumeName(t *testing.T) {
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(filepath.VolumeName(string(p)))
		t1.Result(p.VolumeName())
		t1.AssertEquals()
	}
}
