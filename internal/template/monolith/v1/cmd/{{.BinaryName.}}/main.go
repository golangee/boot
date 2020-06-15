package main

import (
	"flag"
	"github.com/golangee/log"
	zap "github.com/golangee/log-zap"
	_ "{{.Package}}" // reflectplus metadata
	"{{.Package}}/build"
	"{{.Package}}/internal/application"
	"{{.Package}}/internal/config"
	"os"
	"path/filepath"
)

func main() {
	zap.Configure()
	logger := log.New("")
	logger.Info("{{.BinaryName}}", log.Obj("buildCommit", build.Commit), log.Obj("buildTime", build.Time))

	dir, err := os.UserConfigDir()
	if err != nil {
		dir, err = os.UserHomeDir()
		if err != nil {
			dir, err = os.Getwd()
			if err != nil {
				panic(err)
			}
		}
	}

	cfgFile := flag.String("cfg", filepath.Join(dir, config.Filename), "the config file to use")
	help := flag.Bool("help", false, "shows this help")
	frontendDir := flag.String("devFrontend", "", "dev-only: absolute path to a directory with index.html and wasm file")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	app := application.NewServer()
	app.Configure(*cfgFile)
	if *frontendDir != "" {
		logger.Warn("serving development frontend", log.Obj("dir", *frontendDir))
		app.StartDev(*frontendDir)
	}
}
