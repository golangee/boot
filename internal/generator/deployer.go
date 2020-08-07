package generator

import (
	"bytes"
	"fmt"
	"github.com/golangee/boot/internal/resources"
	"github.com/golangee/reflectplus/mod"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var templateableFiles = []string{".go", ".md", "Makefile",".mod"}

type TemplateData struct {
	BinaryName      string
	Path            string
	MainPath        string
	GolangCiVersion string
}

type Options struct {
	Module       *mod.Module
	Dependencies []*mod.Module
}

func Deploy(opts Options) error {
	log.Printf("deploying template into: %s", opts.Module.Dir)
	tplData := &TemplateData{
		BinaryName:      strings.ToLower("myServer"),
		Path:            opts.Module.Path,
		GolangCiVersion: "v1.24.0",
	}

	tplData.MainPath = tplData.Path + "/cmd/" + tplData.BinaryName

	return deployResources(tplData, "monolith/v1", opts.Module.Dir)
}

// deployResources copies and overwrites any files at the destination
func deployResources(tplData *TemplateData, prefix string, dstDir string) error {
	res := resources.Resources()
	for _, file := range res.File {
		if strings.HasPrefix(file.Name, prefix) {
			dstFile := filepath.Join(dstDir, file.Name[len(prefix):])
			tpl, err := template.New(file.Name).Parse(dstFile)
			if err != nil {
				return fmt.Errorf("cannot parse filename '%s': %w", file.Name, err)
			}

			buf := &bytes.Buffer{}
			if err := tpl.Execute(buf, tplData); err != nil {
				return fmt.Errorf("cannot execute filename '%s': %w", file.Name, err)
			}

			dstFile = buf.String()

			if err := os.MkdirAll(filepath.Dir(dstFile), os.ModePerm); err != nil {
				return err
			}

			reader, err := file.Open()
			if err != nil {
				return fmt.Errorf("unable to open zip-file entry '%s': %w", file.Name, err)
			}

			tmp, err := ioutil.ReadAll(reader)
			if err != nil {
				_ = reader.Close()
				return fmt.Errorf("unable to decompress zip-entry '%s': %w", file.Name, err)
			}

			_ = reader.Close()

			needsTemplateProcess := false
			for _, suffix := range templateableFiles {
				if strings.HasSuffix(dstFile, suffix) {
					needsTemplateProcess = true
					break
				}
			}

			if needsTemplateProcess {
				tpl, err := template.New(file.Name).Parse(string(tmp))
				if err != nil {
					return fmt.Errorf("unable to parse file content of '%s': %w", file.Name, err)
				}

				buf.Reset()
				if err := tpl.Execute(buf, tplData); err != nil {
					return fmt.Errorf("cannot execute file content of '%s': %w", file.Name, err)
				}

				if err := ioutil.WriteFile(dstFile, buf.Bytes(), os.ModePerm); err != nil {
					return fmt.Errorf("unable to write templated file: %w", err)
				}
			} else {
				if err := ioutil.WriteFile(dstFile, tmp, os.ModePerm); err != nil {
					return fmt.Errorf("unable to write raw file: %w", err)
				}
			}

			fmt.Println(dstFile)
		}
	}
	return nil
}
