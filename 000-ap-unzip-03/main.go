package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("testdata/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)

		// Open a zip archive for reading.
		r, err := zip.OpenReader(path)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		// Iterate through the files in the archive,
		// printing some of their contents.
		for _, f := range r.File {
			fmt.Println(f.Name)

			// rc, err := f.Open()
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// _, err = io.Copy(os.Stdout, rc)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// rc.Close()
			// fmt.Println()
		}

		return nil
	})
}
