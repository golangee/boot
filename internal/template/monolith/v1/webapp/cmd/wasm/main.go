package main

import (
	"{{.Package}}/internal/application"
)

func main() {
	app := application.NewApp()
	app.Start()
}
