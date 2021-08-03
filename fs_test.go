package pathtype_test

import (
	"io/fs"
	"testing"
)

func TestValidPath(t *testing.T) {
	for _, p := range testPaths {
		res1 := p.ValidPath()
		res2 := fs.ValidPath(string(p))
		if res1 != res2 {
			t.Errorf("path(\"%v\").ValidPath() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)
		}
	}
}

// func TestReadDirFS(t *testing.T) {
// 	for _, p := range testPaths {
// 		res1, err1 := p.ReadDirFS(fsys)
// 		res2, err2 := fs.ReadDir(fsys, string(p))
// 		if !errorsEqual(err1, err2) {
// 			t.Errorf("path(\"%v\").ReadDirFS(%v) errors didn't match. Got: '%v' Expected: '%v'", p, fsys, err1, err2)
// 		}

// 		if len(res1) != len(res2) {
// 			t.Errorf("path(\"%v\").ReadDirFS(%v) result lengths didn't match. Got: '%v' Expected: '%v'", p, fsys, len(res1), len(res2))
// 		} else {
// 			for i, d := range res1 {
// 				if d != res2[i] {
// 					t.Errorf("path(\"%v\").ReadDirFS(%v) didn't match. Got: '%v' Expected: '%v'", p, fsys, res1, res2)
// 				}
// 			}
// 		}

// 	}
// }

// func TestSub(t *testing.T) {
// 	for _, p := range testPaths {
// 		res1, err1 := p.Sub(fsys)
// 		res2, err2 := fs.Sub(fsys, string(p))
// 		if !errorsEqual(err1, err2) {
// 			t.Errorf("path(\"%v\").Sub(%v) errors didn't match. Got: '%v' Expected: '%v'", p, fsys, err1, err2)
// 		}
// 		if res1 != res2 {
// 			t.Errorf("path(\"%v\").Sub(%v) didn't match. Got: '%v' Expected: '%v'", p, fsys, res1, res2)
// 		}
// 	}
// }
