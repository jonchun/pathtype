package pathtype

import (
	"io/fs"
	"os"
	"time"
)

// Executable returns the path name for the executable that started
// the current process. There is no guarantee that the path is still
// pointing to the correct executable. If a symlink was used to start
// the process, depending on the operating system, the result might
// be the symlink or the path it pointed to. If a stable result is
// needed, path/filepath.EvalSymlinks might help.
//
// Executable returns an absolute path unless an error occurred.
//
// The main use case is finding resources located relative to an
// executable.
func Executable() (Path, error) {
	res, err := os.Executable()
	return Path(res), err
}

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func Getwd() (dir Path, err error) {
	res, err := os.Getwd()
	return Path(res), err
}

// TempDir returns the default directory to use for temporary files.
//
// On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
// On Windows, it uses GetTempPath, returning the first non-empty
// value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
// On Plan 9, it returns /tmp.
//
// The directory is neither guaranteed to exist nor have accessible
// permissions.
func TempDir() Path {
	return Path(os.TempDir())
}

// UserCacheDir returns the default root directory to use for user-specific
// cached data. Users should create their own application-specific subdirectory
// within this one and use that.
//
// On Unix systems, it returns $XDG_CACHE_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.cache.
// On Darwin, it returns $HOME/Library/Caches.
// On Windows, it returns %LocalAppData%.
// On Plan 9, it returns $home/lib/cache.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserCacheDir() (Path, error) {
	res, err := os.UserCacheDir()
	return Path(res), err
}

// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserConfigDir() (Path, error) {
	res, err := os.UserConfigDir()
	return Path(res), err
}

// UserHomeDir returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%.
// On Plan 9, it returns the $home environment variable.
func UserHomeDir() (Path, error) {
	res, err := os.UserHomeDir()
	return Path(res), err
}

// Chdir changes the current working directory to the directory at path.
// If there is an error, it will be of type *PathError.
func (path Path) Chdir() error { return os.Chdir(string(path)) }

// Chmod changes the mode of the file at path to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
// If there is an error, it will be of type *PathError.
//
// A different subset of the mode bits are used, depending on the
// operating system.
//
// On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and
// ModeSticky are used.
//
// On Windows, only the 0200 bit (owner writable) of mode is used; it
// controls whether the file's read-only attribute is set or cleared.
// The other bits are currently unused. For compatibility with Go 1.12
// and earlier, use a non-zero mode. Use mode 0400 for a read-only
// file and 0600 for a readable+writable file.
//
// On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive,
// and ModeTemporary are used.
func (path Path) Chmod(mode os.FileMode) error { return os.Chmod(string(path), mode) }

// Chown changes the numeric uid and gid of the file at path.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type *PathError.
//
// On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or
// EPLAN9 error, wrapped in *PathError.
func (path Path) Chown(uid, gid int) error { return os.Chown(string(path), uid, gid) }

// Chtimes changes the access and modification times of the file at path,
// similar to the Unix utime() or utimes() functions.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type *PathError.
func (path Path) Chtimes(atime time.Time, mtime time.Time) error {
	return os.Chtimes(string(path), atime, mtime)
}

// DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory at path.
//
// Note that DirFS("/prefix") only guarantees that the Open calls it makes to the
// operating system will begin with "/prefix": DirFS("/prefix").Open("file") is the
// same as os.Open("/prefix/file"). So if /prefix/file is a symbolic link pointing outside
// the /prefix tree, then using DirFS does not stop the access any more than using
// os.Open does. DirFS is therefore not a general substitute for a chroot-style security
// mechanism when the directory tree contains arbitrary content.
func (path Path) DirFS() fs.FS {
	return os.DirFS(string(path))
}

// Lchown changes the numeric uid and gid of the file at path.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type *PathError.
//
// On Windows, it always returns the syscall.EWINDOWS error, wrapped
// in *PathError.
func (path Path) Lchown(uid, gid int) error {
	return os.Lchown(string(path), uid, gid)
}

// Link creates newname as a hard link to path.
// If there is an error, it will be of type *os.LinkError.
func (path Path) Link(newname string) error {
	return os.Link(string(path), newname)
}

// Mkdir creates a new directory at path with the specified permission
// bits (before umask).
// If there is an error, it will be of type *os.PathError.
func (path Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(path), perm)
}

// MkdirAll creates a directory at path,
// along with any necessary parents, and returns nil,
// or else returns an error.
// The permission bits perm (before umask) are used for all
// directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
func (path Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(path), perm)
}

// MkdirTemp creates a new temporary directory at path
// and returns the pathname of the new directory.
// The new directory's name is generated by adding a random string to the end of pattern.
// If pattern includes a "*", the random string replaces the last "*" instead.
// If path is empty, MkdirTemp uses the default directory for temporary files, as returned by TempDir.
// Multiple programs or goroutines calling MkdirTemp simultaneously will not choose the same directory.
// It is the caller's responsibility to remove the directory when it is no longer needed.
func (path Path) MkdirTemp(pattern string) (Path, error) {
	res, err := os.MkdirTemp(string(path), pattern)
	return Path(res), err
}

// Readlink returns the destination of the symbolic link at path.
// If there is an error, it will be of type *os.PathError.
func (path Path) Readlink() (Path, error) {
	res, err := os.Readlink(string(path))
	return Path(res), err
}

// Remove removes path.
// If there is an error, it will be of type *os.PathError.
func (path Path) Remove() error {
	return os.Remove(string(path))
}

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
// If there is an error, it will be of type *os.PathError.
func (path Path) RemoveAll() error {
	return os.RemoveAll(string(path))
}

// Rename renames (moves) path to newpath.
// If newpath already exists and is not a directory, Rename replaces it.
// OS-specific restrictions may apply when path and newpath are in different directories.
// If there is an error, it will be of type *os.LinkError.
func (path Path) Rename(newpath Path) error {
	return os.Rename(string(path), string(newpath))
}

// Symlink creates newname as a symbolic link to the path.
// If there is an error, it will be of type *os.LinkError..
func (path Path) Symlink(newname string) error {
	return os.Symlink(string(path), newname)
}

// Truncate changes the size of the path.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type *os.PathError.
func (path Path) Truncate(size int64) error {
	return os.Truncate(string(path), size)
}

// WriteFile writes data to the named file, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
func (path Path) WriteFile(data []byte, perm os.FileMode) error {
	return os.WriteFile(string(path), data, perm)
}
