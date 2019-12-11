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
	var src = "log.zip"
	var dst = ""

	if err := UnZip(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func UnZip(dst, src string) (err error) {
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}

			continue
		}

		fr, err := file.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			return err
		}

		fmt.Printf(path, n)

		fw.Close()
		fr.Close()
	}

	return nil
}
