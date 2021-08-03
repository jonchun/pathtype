package pathtype_test

import (
	"fmt"
	"io/fs"
	"path/filepath"

	pt "github.com/jonchun/pathtype"
)

func prepareTestDirTree(tree path) (path, error) {
	tmpDir, err := path("").MkdirTemp("")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = tmpDir.Join(tree).MkdirAll(0755)
	if err != nil {
		tmpDir.RemoveAll()
		return "", err
	}

	return tmpDir, nil
}

func ExampleWalk() {
	type path = pt.Path
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer tmpDir.RemoveAll()
	tmpDir.Chdir()

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = path(".").Walk(func(p path, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", p, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", p)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
	// Output:
	// On Unix:
	// visited file or dir: "."
	// visited file or dir: "dir"
	// visited file or dir: "dir/to"
	// visited file or dir: "dir/to/walk"
	// skipping a dir without errors: skip
}
