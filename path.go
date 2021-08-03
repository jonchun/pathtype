// This is a small wrapper around the filepath API that allows using file paths as their own type
package pathtype

import (
	"io/fs"
	"path/filepath"
)

type Path string

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func (path Path) Abs() (Path, error) {
	res, err := filepath.Abs(string(path))
	return Path(res), err
}

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func (path Path) Base() Path {
	return Path(filepath.Base(string(path)))
}

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
//	1. Replace multiple Separator elements with a single one.
//	2. Eliminate each . path name element (the current directory).
//	3. Eliminate each inner .. path name element (the parent directory)
//	   along with the non-.. element that precedes it.
//	4. Eliminate .. elements that begin a rooted path:
//	   that is, replace "/.." by "/" at the beginning of a path,
//	   assuming Separator is '/'.
//
// The returned path ends in a slash only if it represents a root directory,
// such as "/" on Unix or `C:\` on Windows.
//
// Finally, any occurrences of slash are replaced by Separator.
//
// If the result of this process is an empty string, Clean
// returns the string ".".
//
// See also Rob Pike, ``Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,''
// https://9p.io/sys/doc/lexnames.html
func (path Path) Clean() Path {
	return Path(filepath.Base(string(path)))
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
func (path Path) Dir() Path {
	return Path(filepath.Dir(string(path)))
}

// EvalSymlinks returns path's name after the evaluation of any symbolic
// links.
// If path is relative the result will be relative to the current directory,
// unless one of the components is an absolute symbolic link.
// EvalSymlinks calls Clean on the result.
func (path Path) EvalSymlinks() (Path, error) {
	res, err := filepath.EvalSymlinks(string(path))
	return Path(res), err
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
func (path Path) Ext() string {
	return filepath.Ext(string(path))
}

// FromSlash returns the result of replacing each slash ('/') character
// in path with a separator character. Multiple slashes are replaced
// by multiple separators.
func (path Path) FromSlash() Path {
	return Path(filepath.FromSlash(string(path)))
}

// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in Match. The pattern may describe hierarchical names such as
// /usr/*/bin/ed (assuming the Separator is '/').
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
func (path Path) Glob(pattern string) (matches []Path, err error) {
	p1 := ""
	if filepath.IsAbs(pattern) {
		p1 = pattern
	} else {
		p1 = filepath.Join(string(path), pattern)
	}
	m, err := filepath.Glob(p1)
	if err != nil {
		return
	}
	for _, e := range m {
		matches = append(matches, Path(e))
	}
	return
}

// IsAbs reports whether the path is absolute.
func (path Path) IsAbs() bool {
	return filepath.IsAbs(string(path))
}

// Join joins any number of path elements into path,
// separating them with an OS specific Separator. Empty elements
// are ignored. The result is Cleaned. However, if the argument
// list is empty or all its elements are empty, Join returns
// an empty string.
// On Windows, the result will only be a UNC path if the first
// non-empty element is a UNC path.
func (path Path) Join(elem ...Path) Path {
	var e1 []string
	e1 = append(e1, string(path))
	for _, e := range elem {
		e1 = append(e1, string(e))
	}
	return Path(filepath.Join(e1...))
}

// Match reports whether the path matches the shell file name pattern.
// The pattern syntax is:
//
//	pattern:
//		{ term }
//	term:
//		'*'         matches any sequence of non-Separator characters
//		'?'         matches any single non-Separator character
//		'[' [ '^' ] { character-range } ']'
//		            character class (must be non-empty)
//		c           matches character c (c != '*', '?', '\\', '[')
//		'\\' c      matches character c
//
//	character-range:
//		c           matches character c (c != '\\', '-', ']')
//		'\\' c      matches character c
//		lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
//
// On Windows, escaping is disabled. Instead, '\\' is treated as
// path separator.
//
func (path Path) Match(pattern string) (bool, error) {
	return filepath.Match(pattern, string(path))
}

// Rel returns a relative path that is lexically equivalent to targpath when
// joined to path with an intervening separator. That is,
// path.Join(path.Rel(targpath)) is equivalent to targpath itself.
// On success, the returned path will always be relative to path,
// even if path and targpath share no elements.
// An error is returned if targpath can't be made relative to path or if
// knowing the current working directory would be necessary to compute it.
// Rel calls Clean on the result.
func (path Path) Rel(targpath Path) (Path, error) {
	res, err := filepath.Rel(string(path), string(targpath))
	return Path(res), err
}

// Split splits the path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
func (path Path) Split() (dir, file Path) {
	d, f := filepath.Split(string(path))
	return Path(d), Path(f)
}

// SplitList splits a list of paths joined by the OS-specific ListSeparator,
// usually found in PATH or GOPATH environment variables.
// Unlike strings.Split, SplitList returns an empty slice when passed an empty
// string.
func SplitList(path string) []Path {
	var p1 []Path
	p2 := filepath.SplitList(path)
	for _, ps := range p2 {
		p1 = append(p1, Path(ps))
	}
	return p1
}

// ToSlash returns the result of replacing each separator character
// in path with a slash ('/') character. Multiple separators are
// replaced by multiple slashes.
func (path Path) ToSlash() Path {
	return Path(filepath.ToSlash(string(path)))
}

// VolumeName returns leading volume name.
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share".
// On other platforms it returns "".
func (path Path) VolumeName() Path {
	return Path(filepath.VolumeName(string(path)))
}

// Walk walks the file tree rooted at path, calling fn for each file or
// directory in the tree, including path.
//
// All errors that arise visiting files and directories are filtered by fn:
// see the WalkFunc documentation for details.
//
// The files are walked in lexical order, which makes the output deterministic
// but requires Walk to read an entire directory into memory before proceeding
// to walk that directory.
//
// Walk does not follow symbolic links.
//
// Walk is less efficient than WalkDir, introduced in Go 1.16,
// which avoids calling os.Lstat on every visited file or directory.
func (path Path) Walk(fn filepath.WalkFunc) error {
	return filepath.Walk(string(path), fn)
}

// WalkDir walks the file tree rooted at path, calling fn for each file or
// directory in the tree, including path.
//
// All errors that arise visiting files and directories are filtered by fn:
// see the fs.WalkDirFunc documentation for details.
//
// The files are walked in lexical order, which makes the output deterministic
// but requires WalkDir to read an entire directory into memory before proceeding
// to walk that directory.
//
// WalkDir does not follow symbolic links.
func (path Path) WalkDir(fn fs.WalkDirFunc) error {
	return filepath.WalkDir(string(path), fn)
}
