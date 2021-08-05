package pathtype_test

import (
	"io/fs"
	"os"
	"reflect"
	"testing"
	"testing/fstest"

	pt "github.com/jonchun/pathtype"
)

type path = pt.Path

var fsys = fstest.MapFS{
	"hello world.txt":  {Data: []byte("hi")},
	"hello-world2.txt": {Data: []byte("hola")},
}

var testPaths = []path{
	path("index"),
	path("index.js"),
	path("main.test.js"),
	path("/foo/bar/baz.js"),
	path("/foo/bar/baz"),
	path("/foo/bar/baz/"),
	path("dev.txt"),
	path("../todo.txt"),
	path(".."),
	path("/"),
	path("."),
	path(""),
}

var testFiles = []path{
	path("index"),
	path("index.js"),
	path("main.test.js"),
	path("/foo/bar/baz.js"),
	path("/foo/bar/baz"),
	path("/foo/bar/baz/"),
	path("dev.txt"),
	path("todo.txt"),
}

var testPatterns = []string{
	"abc",
	"*",
	"*c",
	"a*",
	"a*/b",
	"a*b*c*d*e*/f",
	"ab[^b-d]",
	"a\\*b",
}

type TransformFunc func(interface{}) interface{}
type tester struct {
	TB        testing.TB
	Transform TransformFunc
	expect    []interface{}
	result    []interface{}
}

func (t *tester) Expect(val ...interface{}) {
	t.expect = append(t.expect, val...)
}

func (t *tester) Result(val ...interface{}) {
	t.result = append(t.result, val...)
}

// AssertEquals is a helper method to make sure every result is equal to the expected results.
func (t *tester) AssertEquals() {
	t.TB.Helper()
	t.assertLenEqual()
	for i, e := range t.expect {
		r := t.result[i]
		if t.Transform != nil {
			r = t.Transform(r)
		}

		switch e1 := e.(type) {
		case error:
			r1, ok := r.(error)
			if !ok {
				t.Error("expected an error but didn't get one", e1, r1)
				continue
			}
			if e1.Error() != r1.Error() {
				t.Error("errors do not match", e1, r1)
				continue
			}
		case fs.FileInfo:
			r1, ok := r.(fs.FileInfo)
			if !ok {
				t.Error("expected a fs.FileInfo but didn't get one", e1, r1)
				continue
			}
			if !fileInfoEqual(e1, r1) {
				t.Error("File Info not equal", nil, nil)
			}
		case os.File:
			r1, ok := r.(os.File)
			if !ok {
				t.Error("expected a os.File but didn't get one", e1, r1)
				continue
			}
			e2, _ := e1.Stat()
			r2, _ := e1.Stat()
			if !fileInfoEqual(e2, r2) {
				t.Error("File Info not equal", nil, nil)
			}
		case *os.File:
			r1, ok := r.(*os.File)
			if !ok {
				t.Error("expected a *os.File but didn't get one", e1, r1)
				continue
			}
			e2, _ := e1.Stat()
			r2, _ := e1.Stat()
			if !fileInfoEqual(e2, r2) {
				t.Error("File Info not equal", nil, nil)
			}
		default:
			v := reflect.ValueOf(e)
			v1 := reflect.ValueOf(r)
			switch v.Kind() {
			case reflect.Slice, reflect.Array:
				if v.Len() != v1.Len() {
					t.Error("length mismatch", v.Len(), v1.Len())
					continue
				}
				for i := 0; i < v.Len(); i++ {
					a := v.Index(i).Interface()
					b := v1.Index(i).Interface()
					if a != b {
						t.Error("", a, b)
					}
				}
			default:
				if e != r {
					t.Error("", e, r)
				}
			}
		}
	}
}

// Asserts that the types of expected/results are the same and that they are nil/non-nil matching.
func (t *tester) AssertSimilar() {
	t.TB.Helper()
	t.assertLenEqual()
	for i, e := range t.expect {
		r := t.result[i]
		if t.Transform != nil {
			r = t.Transform(r)
		}
		ve := reflect.ValueOf(e)
		vr := reflect.ValueOf(r)
		if ve.IsValid() && vr.IsValid() {
			if ve.IsNil() && !vr.IsNil() {
				t.Error("expected nil", ve, vr)
			} else if !ve.IsNil() && vr.IsNil() {
				t.Error("expected non-nil", ve, vr)
			}
		}
		if ve.Kind() != vr.Kind() {
			t.Error("type mismatch", ve, vr)
		}
	}
}

func (t *tester) assertLenEqual() {
	t.TB.Helper()
	if len(t.expect) != len(t.result) {
		t.Error("number of results do not match", len(t.expect), len(t.result))
		return
	}
}

func fileInfoEqual(f1 fs.FileInfo, f2 fs.FileInfo) bool {
	if f1 != nil && f2 != nil {
		if f1.Name() == f2.Name() ||
			f1.Size() == f2.Size() ||
			f1.Mode() == f2.Mode() ||
			f1.IsDir() == f2.IsDir() {
			return true
		}
	} else if f1 == nil && f2 == nil {
		return true
	}
	return false
}

func (t *tester) Error(msg string, expect, result interface{}) {
	t.TB.Helper()
	if t.TB != nil {
		t.TB.Errorf("%s: expect %v | result %v", msg, expect, result)
	} else {
		panic("FATAL NO TB")
	}
}

var pathToString TransformFunc = func(in interface{}) interface{} {
	switch in1 := in.(type) {
	case path:
		return string(in1)
	case []path:
		var out []string
		for _, p := range in1 {
			out = append(out, string(p))
		}
		return out
	default:
		return in
	}
}
