package pathtype_test

import (
	"bytes"
	"io/fs"
	"os"
	"strings"
	"testing"
	"time"

	pt "github.com/jonchun/pathtype"
)

func TestExecutable(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.Executable())
	t1.Result(pt.Executable())
	t1.AssertEquals()
}

func TestGetwd(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.Getwd())
	t1.Result(pt.Getwd())
	t1.AssertEquals()
}

func TestTempDir(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.TempDir())
	t1.Result(pt.TempDir())
	t1.AssertEquals()
}

func TestUserCacheDir(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.UserCacheDir())
	t1.Result(pt.UserCacheDir())
	t1.AssertEquals()
}

func TestUserConfigDir(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.UserConfigDir())
	t1.Result(pt.UserConfigDir())
	t1.AssertEquals()
}

func TestUserHomeDir(t *testing.T) {
	t1 := tester{TB: t, Transform: pathToString}
	t1.Expect(os.UserHomeDir())
	t1.Result(pt.UserHomeDir())
	t1.AssertEquals()
}

func TestChdir(t *testing.T) {
	oldD, _ := pt.Getwd()
	defer oldD.Chdir()

	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Chdir(string(p)))
		t1.Result(p.Chdir())
		t1.AssertEquals()
	}
}

func TestChmod(t *testing.T) {
	modes := []fs.FileMode{0644, 0755, 0777}
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testFiles {
		for _, mode := range modes {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Chmod(string(p), mode))
			t1.Result(p.Chmod(mode))
			t1.AssertEquals()
		}
	}
}

func TestChown(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()

	for _, paths := range [][]path{testPaths, testFiles} {
		for _, p := range paths {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Chown(string(p), os.Geteuid(), os.Getegid()))
			t1.Result(p.Chown(os.Geteuid(), os.Getegid()))
			t1.AssertEquals()

			t2 := tester{TB: t, Transform: pathToString}
			t2.Expect(os.Chown(string(p), 0, 0))
			t2.Result(p.Chown(0, 0))
			t2.AssertEquals()
		}
	}
}

func TestChtimes(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()

	for _, paths := range [][]path{testPaths, testFiles} {
		for _, p := range paths {
			now := time.Now()
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Chtimes(string(p), now, now))
			t1.Result(p.Chtimes(now, now))
			t1.AssertEquals()
		}
	}
}

func TestCreate(t *testing.T) {
	d, err := path("").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	oldD, _ := pt.Getwd()
	defer oldD.Chdir()

	d.Chdir()

	for _, paths := range [][]path{testPaths, testFiles} {
		for _, p := range paths {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Create(string(p)))
			t1.Result(p.Create())
			t1.AssertEquals()
		}
	}
}

func TestCreateTemp(t *testing.T) {
	d, err := path("").MkdirTemp("temp*")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	for _, p := range testFiles {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.CreateTemp(string(d), string(p)))
		t1.Result(d.CreateTemp(string(p)))
		t1.AssertSimilar()
	}
}

func TestLchown(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()

	for _, paths := range [][]path{testPaths, testFiles} {
		for _, p := range paths {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Lchown(string(p), os.Geteuid(), os.Getegid()))
			t1.Result(p.Lchown(os.Geteuid(), os.Getegid()))
			t1.AssertEquals()

			t2 := tester{TB: t, Transform: pathToString}
			t2.Expect(os.Lchown(string(p), 0, 0))
			t2.Result(p.Lchown(0, 0))
			t2.AssertEquals()
		}
	}
}

func TestLink(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		for _, p1 := range testPaths {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Link(string(p), string(p1)))
			p1.Remove()
			t1.Result(p.Link(p1))
			p1.Remove()
			t1.AssertEquals()
		}
	}
}

func TestLstat(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Lstat(string(p)))
		t1.Result(p.Lstat())
		t1.AssertEquals()
	}
}

func TestMkdir(t *testing.T) {
	d, err := path(".").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	err = d.Mkdir(0755)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		t.Errorf("path(\"%v\").Mkdir() failed. error: `%v`", d, err)
	}
	err = d.Remove()
	if err != nil {
		panic(err)
	}
	err = d.Mkdir(0755)
	if err != nil {
		t.Errorf("path(\"%v\").Mkdir() failed. error: `%v`", d, err)
	}
}

func TestMkdirAll(t *testing.T) {
	d, err := path(".").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	err = d.MkdirAll(0755)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		t.Errorf("path(\"%v\").MkdirAll() failed. error: `%v`", d, err)
	}
	err = d.Remove()
	if err != nil {
		panic(err)
	}
	err = d.MkdirAll(0755)
	if err != nil {
		t.Errorf("path(\"%v\").MkdirAll() failed. error: `%v`", d, err)
	}
}

func TestMkdirTemp(t *testing.T) {
	a := "."
	d, err := path(a).MkdirTemp("")
	if err != nil {
		t.Errorf("path(\"%v\").MkdirTemp(\"\") failed. error: `%v`", a, err)
	}
	defer d.RemoveAll()
}

func TestOpen(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Open(string(p)))
		t1.Result(p.Open())
		t1.AssertEquals()
	}
	for _, p := range testFiles {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Open(string(p)))
		t1.Result(p.Open())
		t1.AssertEquals()
	}
}

func TestOpenFile(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.OpenFile(string(p), os.O_RDWR, 0))
		t1.Result(p.OpenFile(os.O_RDWR, 0))
		t1.AssertEquals()
	}
	for _, p := range testFiles {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.OpenFile(string(p), os.O_RDWR, 0))
		t1.Result(p.OpenFile(os.O_RDWR, 0))
		t1.AssertEquals()
	}
}

func TestReadlink(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testPaths {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Readlink(string(p)))
		t1.Result(p.Readlink())
		t1.AssertEquals()
	}
}

func TestRemove(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testFiles {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Remove(string(p)))
		t1.Result(p.Remove())
		t1.AssertEquals()
	}
}

func TestRemoveAll(t *testing.T) {
	d := createFilesInTmp(testFiles)
	err := d.RemoveAll()
	if err != nil {
		t.Errorf("path(\"%v\").RemoveAll() failed. error: `%v`", d, err)
	}
}

func TestRename(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testFiles {
		renameTo := p.Join(path("test"))
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Rename(string(p), string(renameTo)))
		t1.Result(p.Rename(renameTo))
		t1.AssertEquals()
	}
}

func TestStat(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testFiles {
		t1 := tester{TB: t, Transform: pathToString}
		t1.Expect(os.Stat(string(p)))
		t1.Result(p.Stat())
		t1.AssertEquals()
	}
}

func TestSymlink(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		for _, p1 := range testFiles {
			t1 := tester{TB: t, Transform: pathToString}
			t1.Expect(os.Symlink(string(p), string(p1)))
			p1.Remove()
			t1.Result(p.Symlink(p1))
			p1.Remove()
			t1.AssertEquals()
		}
	}
}

func TestTruncate(t *testing.T) {
	d := createFilesInTmp(testFiles)
	defer d.RemoveAll()
	for _, p := range testFiles {
		s, err := p.Stat()
		if err != nil {
			if !strings.Contains(err.Error(), "no such file or directory") {
				t.Error(err)
			}
			continue
		}
		// oldSize := s.Size()
		var truncVal int64 = 1234
		err = p.Truncate(truncVal)
		if err != nil {
			t.Errorf("path(\"%v\").Truncate(\"%v\") errored: %v", p, truncVal, err)
			return
		}
		if s.Size() != truncVal {
			t.Errorf("path(\"%v\").Truncate(\"%v\") did not give expected size. Got %v", p, truncVal, s.Size())
		}
	}
}

func TestWriteFile(t *testing.T) {
	d, err := path(".").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()
	f, err := d.CreateTemp("")
	if err != nil {
		panic(err)
	}
	p := path(f.Name())
	mode := fs.FileMode(0644)
	buf := new(bytes.Buffer)
	buf.Grow(10)
	testString := "123"
	buf.WriteString(testString)
	err = p.WriteFile(buf.Bytes(), mode)
	if err != nil {
		t.Errorf("path(\"%v\").WriteFile(\"%v\", \"%v\") failed. \nerror: %v", p, buf.Bytes(), mode, err)
	}
	f, err = p.Open()
	if err != nil {
		panic(err)
	}
	bSlice := make([]byte, len(testString))
	_, err = f.Read(bSlice)
	if err != nil {
		panic(err)
	}
	if string(bSlice) != testString {
		t.Errorf("path(\"%v\").WriteFile(\"%v\", \"%v\") failed. Got: '%v' Expected: '%v'", p, buf.Bytes(), mode, string(bSlice), testString)
	}
}

func createFilesInTmp(paths []path) path {
	oldD, _ := pt.Getwd()
	defer oldD.Chdir()
	d, err := path(".").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	d.Chdir()

	for _, p := range paths {
		p.Create()
	}
	return d
}
