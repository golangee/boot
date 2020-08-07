//go:generate go run embedd.go
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/golangee/boot/internal/util"
	"github.com/golangee/reflectplus/mod"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	modules, err := mod.List(dir)
	if err != nil {
		return err
	}

	if len(modules) == 0 || !modules[0].Main {
		return fmt.Errorf("no main module")
	}

	srcDir := filepath.Join(modules[0].Dir, "internal", "template")
	buf := &bytes.Buffer{}
	if err := util.ZipDir(srcDir, buf); err != nil {
		return err
	}

	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	dstFile := filepath.Join(modules[0].Dir, "internal", "resources", "data.go")
	sb := &bytes.Buffer{}
	sb.WriteString("package resources\n\n")
	sb.WriteString("var data = \"" + b64 + "\"\n")

	if err := ioutil.WriteFile(dstFile, sb.Bytes(), os.ModePerm); err != nil {
		return err
	}

	log.Println("written to", dstFile)
	return nil
}
