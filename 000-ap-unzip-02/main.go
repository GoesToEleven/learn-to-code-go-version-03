package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var counter int
	filepath.Walk("testdata/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		counter++
		fmt.Println(path)
		return nil
	})
	fmt.Println(counter)

	// // Open a zip archive for reading.
	// r, err := zip.OpenReader("testdata/157-custom-views.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer r.Close()

	// // Iterate through the files in the archive,
	// // printing some of their contents.
	// for _, f := range r.File {
	// 	fmt.Printf("Contents of %s:\n", f.Name)
	// 	rc, err := f.Open()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = io.CopyN(os.Stdout, rc, 68)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rc.Close()
	// 	fmt.Println()
	// }
}
