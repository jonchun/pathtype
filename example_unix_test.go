package pathtype_test

import (
	"fmt"

	"github.com/jonchun/pathtype"
)

func ExampleSplitList() {
	fmt.Println("On Unix:", pathtype.SplitList("/a/b/c:/usr/bin"))
	// Output:
	// On Unix: [/a/b/c /usr/bin]
}

func ExamplePath_Rel() {
	paths := []path{
		path("/a/b/c"),
		path("/b/c"),
		path("./b/c"),
	}
	base := path("/a")

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := base.Rel(p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

	// Output:
	// On Unix:
	// "/a/b/c": "b/c" <nil>
	// "/b/c": "../b/c" <nil>
	// "./b/c": "" Rel: can't make ./b/c relative to /a
}

func ExamplePath_Split() {
	paths := []path{
		path("/home/arnie/amelia.jpg"),
		path("/mnt/photos/"),
		path("rabbit.jpg"),
		path("/usr/local//go"),
	}
	fmt.Println("On Unix:")
	for _, p := range paths {
		dir, file := p.Split()
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
	// Output:
	// On Unix:
	// input: "/home/arnie/amelia.jpg"
	// 	dir: "/home/arnie/"
	// 	file: "amelia.jpg"
	// input: "/mnt/photos/"
	// 	dir: "/mnt/photos/"
	// 	file: ""
	// input: "rabbit.jpg"
	// 	dir: ""
	// 	file: "rabbit.jpg"
	// input: "/usr/local//go"
	// 	dir: "/usr/local//"
	// 	file: "go"
}

func ExamplePath_Join() {
	fmt.Println("On Unix:")
	fmt.Println(path("a").Join(path("b"), path("c")))
	fmt.Println(path("a").Join(path("b/c")))
	fmt.Println(path("a/b").Join(path("c")))
	fmt.Println(path("a/b").Join(path("/c")))

	fmt.Println(path("a/b").Join(path("../../../xyz")))

	// Output:
	// On Unix:
	// a/b/c
	// a/b/c
	// a/b/c
	// a/b/c
	// ../xyz
}

func ExamplePath_Match() {
	fmt.Println("On Unix:")
	fmt.Println(path("/home/catch/foo").Match("/home/catch/*"))
	fmt.Println(path("/home/catch/foo/bar").Match("/home/catch/*"))
	fmt.Println(path("/home/gopher").Match("/home/?opher"))
	fmt.Println(path("/home/*").Match("/home/\\*"))

	// Output:
	// On Unix:
	// true <nil>
	// false <nil>
	// true <nil>
	// true <nil>
}

func ExamplePath_Base() {
	fmt.Println("On Unix:")
	fmt.Println(path("/foo/bar/baz.js").Base())
	fmt.Println(path("/foo/bar/baz").Base())
	fmt.Println(path("/foo/bar/baz/").Base())
	fmt.Println(path("dev.txt").Base())
	fmt.Println(path("../todo.txt").Base())
	fmt.Println(path("..").Base())
	fmt.Println(path(".").Base())
	fmt.Println(path("/").Base())
	fmt.Println(path("").Base())

	// Output:
	// On Unix:
	// baz.js
	// baz
	// baz
	// dev.txt
	// todo.txt
	// ..
	// .
	// /
	// .
}
func ExamplePath_Dir() {
	fmt.Println("On Unix:")
	fmt.Println(path("/foo/bar/baz.js").Dir())
	fmt.Println(path("/foo/bar/baz").Dir())
	fmt.Println(path("/foo/bar/baz/").Dir())
	fmt.Println(path("/dirty//path///").Dir())
	fmt.Println(path("dev.txt").Dir())
	fmt.Println(path("../todo.txt").Dir())
	fmt.Println(path("..").Dir())
	fmt.Println(path(".").Dir())
	fmt.Println(path("/").Dir())
	fmt.Println(path("").Dir())

	// Output:
	// On Unix:
	// /foo/bar
	// /foo/bar
	// /foo/bar/baz
	// /dirty/path
	// .
	// ..
	// .
	// .
	// /
	// .
}

func ExamplePath_IsAbs() {
	fmt.Println("On Unix:")
	fmt.Println(path("/home/gopher").IsAbs())
	fmt.Println(path(".bashrc").IsAbs())
	fmt.Println(path("..").IsAbs())
	fmt.Println(path(".").IsAbs())
	fmt.Println(path("/").IsAbs())
	fmt.Println(path("").IsAbs())

	// Output:
	// On Unix:
	// true
	// false
	// false
	// false
	// true
	// false
}
