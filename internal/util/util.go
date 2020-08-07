package util

import (
	"archive/zip"
	. "github.com/golangee/boot/internal/errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ZipDir(srcDir string, dst io.Writer) (err error) {
	writer := zip.NewWriter(dst)
	defer Check(writer.Close, &err)
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) (e error) {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer Check(file.Close, &e)

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		log.Print(relPath)
		f, err := writer.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil

	})

	return
}
