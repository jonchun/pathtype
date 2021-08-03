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
	res1, err1 := pt.Executable()
	res2, err2 := os.Executable()
	if !errorsEqual(err1, err2) {
		t.Errorf("pt.Executable() errors didn't match. Got: '%v' Expected: '%v'", err1, err2)
	}
	if string(res1) != res2 {
		t.Errorf("pt.Executable() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestGetwd(t *testing.T) {
	res1, err1 := pt.Getwd()
	res2, err2 := os.Getwd()
	if !errorsEqual(err1, err2) {
		t.Errorf("pt.Getwd() errors didn't match. Got: '%v' Expected: '%v'", err1, err2)
	}
	if string(res1) != res2 {
		t.Errorf("pt.Getwd() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestTempDir(t *testing.T) {
	res1 := pt.TempDir()
	res2 := os.TempDir()
	if string(res1) != res2 {
		t.Errorf("pt.TempDir() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestUserCacheDir(t *testing.T) {
	res1, err1 := pt.UserCacheDir()
	res2, err2 := os.UserCacheDir()
	if !errorsEqual(err1, err2) {
		t.Errorf("pt.UserCacheDir() errors didn't match. Got: '%v' Expected: '%v'", err1, err2)
	}
	if string(res1) != res2 {
		t.Errorf("pt.UserCacheDir() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestUserConfigDir(t *testing.T) {
	res1, err1 := pt.UserConfigDir()
	res2, err2 := os.UserConfigDir()
	if !errorsEqual(err1, err2) {
		t.Errorf("pt.UserConfigDir() errors didn't match. Got: '%v' Expected: '%v'", err1, err2)
	}
	if string(res1) != res2 {
		t.Errorf("pt.UserConfigDir() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestUserHomeDir(t *testing.T) {
	res1, err1 := pt.UserHomeDir()
	res2, err2 := os.UserHomeDir()
	if !errorsEqual(err1, err2) {
		t.Errorf("pt.UserHomeDir() errors didn't match. Got: '%v' Expected: '%v'", err1, err2)
	}
	if string(res1) != res2 {
		t.Errorf("pt.UserHomeDir() didn't match. Got: '%v' Expected: '%v'", res1, res2)
	}
}

func TestChdir(t *testing.T) {
	oldD, _ := pt.Getwd()
	defer oldD.Chdir()
	for _, p := range testPaths {
		err1 := p.Chdir()
		err2 := os.Chdir(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Chdir() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
	}
}

func TestChmod(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		mode := fs.FileMode(0644)
		err1 := p.Chmod(mode)
		err2 := os.Chmod(string(p), mode)
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Chmod(%v) errors didn't match. Got: '%v' Expected: '%v'", p, mode, err1, err2)
		}
	}
}

func TestChown(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		err := p.Chown(os.Geteuid(), os.Getegid())
		if err != nil {
			t.Errorf("path(\"%v\").Chown(%v, %v) failed", p, os.Geteuid(), os.Getegid())
		}
		err = p.Chown(0, 0)
		if err != nil && !strings.Contains(err.Error(), "operation not permitted") {
			t.Errorf("path(\"%v\").Chown(%v, %v) failed", p, 0, 0)
		}
	}
}

func TestChtimes(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		now := time.Now()
		err1 := p.Chtimes(now, now)
		err2 := os.Chtimes(string(p), now, now)
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Chtimes(%v, %v) errors didn't match. Got: '%v' Expected: '%v'", p, now, now, err1, err2)
		}
	}
}
func TestCreate(t *testing.T) {
	d, err := path(".").MkdirTemp("")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	oldD, _ := pt.Getwd()
	defer oldD.Chdir()
	d.Chdir()
	for _, p := range testPaths2 {
		var name1, name2 string
		file1, err1 := p.Create()
		if file1 != nil {
			name1 = file1.Name()
		}
		path(name1).Remove()
		file2, err2 := os.Create(string(p))
		if file2 != nil {
			name2 = file2.Name()
		}
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Create() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if name1 != name2 {
			t.Errorf("path(\"%v\").Create() didn't match. Got: '%v' Expected: '%v'", p, name1, name2)
		}
	}
}

func TestCreateTemp(t *testing.T) {
	d, err := path(".").MkdirTemp("temp*")
	if err != nil {
		panic(err)
	}
	defer d.RemoveAll()

	_, err = d.CreateTemp("hello*")
	if err != nil {
		t.Errorf("path(\"%v\").CreateTemp() failed. error: `%v`", d, err)
	}
}

func TestLchown(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		err := p.Lchown(os.Geteuid(), os.Getegid())
		if err != nil {
			t.Errorf("path(\"%v\").Lchown(%v, %v) failed", p, os.Geteuid(), os.Getegid())
		}
		err = p.Lchown(0, 0)
		if err != nil && !strings.Contains(err.Error(), "operation not permitted") {
			t.Errorf("path(\"%v\").Lchown(%v, %v) failed", p, 0, 0)
		}
	}
}

func TestLink(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		for _, p1 := range testPaths2 {
			err1 := p.Link(p1)
			p1.Remove()
			err2 := os.Link(string(p), string(p1))
			p1.Remove()
			if !errorsEqual(err1, err2) {
				t.Errorf("path(\"%v\").Link(\"%v\") errors didn't match. Got: '%v' Expected: '%v'", p, p1, err1, err2)
			}
		}
	}
}

func TestLstat(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		file1, err1 := p.Lstat()
		file2, err2 := os.Lstat(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Lstat() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if file1 != nil && file2 != nil {
			if !(file1.Name() == file2.Name() && file1.Size() == file2.Size()) {
				t.Errorf("path(\"%v\").Lstat() didn't match. Got: '%v' Expected: '%v'", p, file1, file2)
			}
		}
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
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		_, err := p.Open()
		if err != nil {
			t.Errorf("path(\"%v\").Open() failed. error: `%v`", p, err)
		}
	}
	for _, p := range testPaths2 {
		_, err := p.Open()
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			t.Errorf("path(\"%v\").Open() failed. error: `%v`", p, err)
		}
	}
}

func TestOpenFile(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		_, err := p.OpenFile(os.O_RDONLY, 0)
		if err != nil {
			t.Errorf("path(\"%v\").OpenFile() failed. error: `%v`", p, err)
		}
	}
	for _, p := range testPaths2 {
		_, err := p.Open()
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			t.Errorf("path(\"%v\").OpenFile() failed. error: `%v`", p, err)
		}
	}
}

func TestReadlink(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		res1, err1 := p.Readlink()
		res2, err2 := os.Readlink(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Readlink() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if string(res1) != res2 {
			t.Errorf("path(\"%v\").Readlink() didn't match. Got: '%v' Expected: '%v'", p, res1, res2)

		}
	}
}

func TestRemove(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		err := p.Remove()
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			t.Errorf("path(\"%v\").Remove() failed. error: `%v`", p, err)
		}
	}
}

func TestRemoveAll(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	err := d.RemoveAll()
	if err != nil {
		t.Errorf("path(\"%v\").RemoveAll() failed. error: `%v`", d, err)
	}
}

func TestRename(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		err := p.Rename(p.Join(path("test")))
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			t.Errorf("path(\"%v\").Rename(\"%v\") failed. error: `%v`", p, p.Join(path("test")), err)
		}
	}
}

func TestStat(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
		file1, err1 := p.Stat()
		file2, err2 := os.Stat(string(p))
		if !errorsEqual(err1, err2) {
			t.Errorf("path(\"%v\").Stat() errors didn't match. Got: '%v' Expected: '%v'", p, err1, err2)
		}
		if file1 != nil && file2 != nil {
			if !(file1.Name() == file2.Name() && file1.Size() == file2.Size()) {
				t.Errorf("path(\"%v\").Stat() didn't match. Got: '%v' Expected: '%v'", p, file1, file2)
			}
		}
	}
}

func TestSymlink(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	globs, _ := d.Glob("*")
	for _, p := range globs {
		for _, p1 := range testPaths2 {
			err1 := p.Symlink(p1)
			p1.Remove()
			err2 := os.Symlink(string(p), string(p1))
			p1.Remove()
			if !errorsEqual(err1, err2) {
				t.Errorf("path(\"%v\").Symlink(\"%v\") errors didn't match. Got: '%v' Expected: '%v'", p, p1, err1, err2)
			}
		}
	}
}

func TestTruncate(t *testing.T) {
	d := createFilesInTmp(testPaths2)
	defer d.RemoveAll()
	for _, p := range testPaths2 {
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
