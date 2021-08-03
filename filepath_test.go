package pathtype_test

import (
	"path/filepath"
	"sort"
	"testing"
)

func TestAbs(t *testing.T) {
	for _, p := range testPaths {
		res1, err1 := p.Abs()
		res2, err2 := filepath.Abs(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Abs() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").Abs() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestBase(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.Base()
		res2 := filepath.Base(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").Base() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestClean(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.Clean()
		res2 := filepath.Clean(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").Clean() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestDir(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.Dir()
		res2 := filepath.Dir(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").Dir() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestEvalSymlinks(t *testing.T) {
	for _, p := range testPaths {
		res1, err1 := p.EvalSymlinks()
		res2, err2 := filepath.EvalSymlinks(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").EvalSymlinks() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").EvalSymlinks() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestExt(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.Ext()
		res2 := filepath.Ext(string(p))
		if res1 != res2 {
			t.Errorf("path(\"%v\").Ext() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestFromSlash(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.FromSlash()
		res2 := filepath.FromSlash(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").FromSlash() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestGlob(t *testing.T) {
	for _, p := range testPaths {
		for _, p1 := range testPatterns {
			res1, err1 := p.Glob(p1)
			res2, err2 := filepath.Glob(filepath.Join(string(p), p1))
			if !errorsEqual(err1, err2) {
				t.Errorf("path(\"%v\").Glob(\"%s\") errors didn't match. Got: '%v' Expected: '%v'", p, p1, err1, err2)
			}

			// convert res1 to a slice of strings
			var res3 []string
			for _, m := range res1 {
				res3 = append(res3, string(m))
			}

			// sort both (unnecessary?)
			sort.Strings(res2)
			sort.Strings(res3)

			for i, m := range res3 {
				if m != res2[i] {
					t.Errorf("path(\"%v\").Glob(\"%s\") didn't match. Got: '%v' Expected: '%v'", p, p1, m, res2[i])
				}
			}
		}
	}
}

func TestIsAbs(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.IsAbs()
		res2 := filepath.IsAbs(string(p))
		if res1 != res2 {
			t.Errorf("path(\"%v\").IsAbs() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

func TestMatch(t *testing.T) {
	for _, p := range testPaths {
		for _, p1 := range testPatterns {
			res1, err1 := p.Match(p1)
			res2, err2 := filepath.Match(p1, string(p))
			if !errorsEqual(err1, err2) {
				t.Errorf("path(\"%v\").Match(\"%s\") errors didn't match. Got: '%v' Expected: '%v'", p, p1, err1, err2)
			}
			if res1 != res2 {
				t.Errorf("path(\"%v\").Match() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
			}
		}
	}
}

func TestRel(t *testing.T) {
	for _, p := range testPaths {
		for _, targpath := range testPaths {
			res1, err1 := p.Rel(targpath)
			res2, err2 := filepath.Rel(string(p), string(targpath))
			if !errorsEqual(err1, err2) {
				t.Errorf("path(\"%v\").Rel(\"%v\") errors didn't match. Got: '%v' Expected: '%v'", p, targpath, err1, err2)
			}
			if string(res1) != res2 {
				t.Errorf("path(\"%v\").Rel(\"%v\") didn't match. Got: '%v' Expected: '%v'", p, targpath, res1, res2)
			}
		}
	}
}

func TestSplit(t *testing.T) {
	for _, p := range testPaths {
		dir1, file1 := p.Split()
		dir2, file2 := filepath.Split(string(p))
		if string(dir1) != dir2 {
			t.Errorf("path(\"%v\").Split() dir didn't match. Got: '%v' Expected: '%v'", p, dir1, dir2)
		}
		if string(file1) != file2 {
			t.Errorf("path(\"%v\").Split() file didn't match. Got: '%v' Expected: '%v'", p, file1, file2)
		}
	}
}

func TestToSlash(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.ToSlash()
		res2 := filepath.ToSlash(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").ToSlash() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}
func TestVolumeName(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.VolumeName()
		res2 := filepath.VolumeName(string(p))
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").VolumeName() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}
