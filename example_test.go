package pathtype_test

import (
	"fmt"

	"github.com/jonchun/pathtype"
)

type path = pathtype.Path

func ExamplePath_Ext() {
	fmt.Printf("No dots: %q\n", path("index").Ext())
	fmt.Printf("One dot: %q\n", path("index.js").Ext())
	fmt.Printf("Two dots: %q\n", path("main.test.js").Ext())
	// Output:
	// No dots: ""
	// One dot: ".js"
	// Two dots: ".js"
}
