package pathtype

import (
	"os"
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
