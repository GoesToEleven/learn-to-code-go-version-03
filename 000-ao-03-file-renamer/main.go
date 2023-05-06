package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Name() == "renamer" || info.Name() == "." {
			fmt.Printf("skipped without errors: %+v \n", info.Name())
			return nil
		}
		fmt.Printf("visited file or dir: %q\n", path)
		xs := strings.Split(path, ".")
		name := xs[0]
		ext := "NOTHING"
		if len(xs) == 2 {			
			ext = xs[1]
		}
		newPath := name + "XXX" + "." + ext
		os.Rename(path, newPath)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}
}