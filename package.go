// Path boot contains a template system to bootstrap a golangee opionated setup.
package boot

import (
	"fmt"
	"github.com/golangee/boot/internal/generator"
	"github.com/golangee/reflectplus/mod"
	"os"
	"path/filepath"
)

func MustGenerate() {
	if err := Generate(); err != nil {
		panic(err)
	}
}

func Generate() error {
	return GenerateWithOptions(false)
}

func GenerateWithOptions(forceDeployment bool) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	modules, err := mod.List(dir)
	if err != nil {
		return fmt.Errorf("failed to list modules:%w", err)
	}

	if len(modules) == 0 {
		return fmt.Errorf("no modules available")
	}

	if !modules[0].Main {
		return fmt.Errorf("no main module available")
	}

	modRootDir := modules[0].Dir
	_, err = os.Stat(filepath.Join(modRootDir, "Makefile"))
	needStubGeneration := err != nil || forceDeployment

	if needStubGeneration {
		if err := generator.Deploy(generator.Options{
			Module:       modules[0],
			Dependencies: modules[1:],
		}); err != nil {
			return fmt.Errorf("found empty module but cannot deploy stub: %w", err)
		}
	}

	return nil
}
