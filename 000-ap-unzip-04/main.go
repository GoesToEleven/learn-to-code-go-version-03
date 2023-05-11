package main

import (
	"archive/zip"
	"fmt"
	"io"
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

			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer rc.Close()

			nf, err := os.Create(f.Name)
			if err != nil {
				log.Fatal(err)
			}
			defer nf.Close()

			_, err = io.Copy(nf, rc)
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})
}
