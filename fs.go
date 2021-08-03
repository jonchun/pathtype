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

// ReadDir reads the directory at path
// and returns a list of directory entries sorted by filename.
//
// If fs implements ReadDirFS, ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func (path Path) ReadDir(fsys fs.FS) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, string(path))
}

// Sub returns an FS corresponding to the subtree rooted at fsys's directory located at path.
//
// If fs implements SubFS, Sub calls returns fsys.Sub(path).
// Otherwise, if dir is ".", Sub returns fsys unchanged.
// Otherwise, Sub returns a new FS implementation sub that,
// in effect, implements sub.Open(path) as fsys.Open(path.Join(path, name)).
// The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.
//
// Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix")
// and that neither of them guarantees to avoid operating system
// accesses outside "/prefix", because the implementation of os.DirFS
// does not check for symbolic links inside "/prefix" that point to
// other directories. That is, os.DirFS is not a general substitute for a
// chroot-style security mechanism, and Sub does not change that fact.
func (path Path) Sub(fsys fs.FS) (fs.FS, error) {
	return fs.Sub(fsys, string(path))
}

// Stat returns a FileInfo describing the file at path from the file system.
//
// If fs implements StatFS, Stat calls fs.Stat.
// Otherwise, Stat opens the file to stat it.
func (path Path) Stat(fsys fs.FS) (fs.FileInfo, error) {
	return fs.Stat(fsys, string(path))
}
