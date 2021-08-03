package pathtype_test

import (
	pt "github.com/jonchun/pathtype"
)

type path = pt.Path

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

var testPaths2 = []path{
	path("index"),
	path("index.js"),
	path("main.test.js"),
	path("/foo/bar/baz.js"),
	path("/foo/bar/baz"),
	path("/foo/bar/baz/"),
	path("dev.txt"),
	path("../todo.txt"),
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

// Checks to see if two errors are equal
func errorsEqual(err1 error, err2 error) bool {
	isEqual := false
	if err1 == nil && err2 == nil {
		isEqual = true
	} else if err1 != nil && err2 != nil {
		isEqual = err1.Error() == err2.Error()
	}
	return isEqual
}
