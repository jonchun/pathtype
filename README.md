# pathtype

Add a type for paths in Go. This small package basically just wraps "path/filepath" from the Standard library.

## Example

### Code

```
package main

import (
	"fmt"
	"log"

	pt "github.com/jonchun/pathtype"
)

type path = pt.Path

func main() {
	myFile := path("myfile.txt")
	exampleFile := path("example/example.txt")
	fmt.Println(exampleFile.Dir())
	fmt.Println(exampleFile.Dir().Join(myFile))

	res, err := exampleFile.Dir().Join(myFile).Dir().Abs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	fmt.Println("=========================")
	listBase(res)
	fmt.Println("=========================")
	listExt(res)
}

// list all Base for files in p
func listBase(p path) {
	if glob, err := p.Glob("*"); err != nil {
		log.Fatal(err)
	} else {
		for _, match := range glob {
			fmt.Println(match.Base())
		}
	}
}

// list all extensions for files in p
func listExt(p path) {
	if glob, err := p.Glob("*"); err != nil {
		log.Fatal(err)
	} else {
		for _, match := range glob {
			fmt.Println(match.Ext())
		}
	}
}
```

### Output

```
example
example/myfile.txt
/home/jonchun/example_module/example
=========================
1.log
2.log
example.txt
=========================
.log
.log
.txt
```

See [GoDoc](https://godoc.org/github.com/jonchun/pathtype) for documentation, but it should be pretty self-explanatory.

## TODO

- Add wrappers for other packages that take paths as strings. e.g: [os](https://pkg.go.dev/os)
  Would be nice to have syntax similar to

  ```
  import pt "github.com/jonchun/pathtype"
  type path = pt.Path

  func example(p path) {
      pt.Chmod(p, 0644)
  }
  ```
