package pathtype

import "io/fs"

// ValidPath reports whether the given path
// is valid for use in a call to Open.
//
// Path names passed to open are UTF-8-encoded,
// unrooted, slash-separated sequences of path elements, like “x/y/z”.
// Path names must not contain an element that is “.” or “..” or the empty string,
// except for the special case that the root directory is named “.”.
// Paths must not start or end with a slash: “/x” and “x/” are invalid.
//
// Note that paths are slash-separated on all systems, even Windows.
// Paths containing other characters such as backslash and colon
// are accepted as valid, but those characters must never be
// interpreted by an FS implementation as path element separators.
func (path Path) ValidPath() bool {
	return fs.ValidPath(string(path))
}

// ReadDirFS reads the directory at path in FS fsys.
// and returns a list of directory entries sorted by filename.
func (path Path) ReadDirFS(fsys fs.FS) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, string(path))
}

// Sub returns an FS corresponding to the subtree rooted at fsys's directory located at path.
func (path Path) Sub(fsys fs.FS) (fs.FS, error) {
	return fs.Sub(fsys, string(path))
}

// WalkDirFunc is the type of the function called by WalkDir to visit each file or directory.
type WalkDirFunc func(path Path, d fs.DirEntry, err error) error
