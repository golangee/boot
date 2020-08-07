package main

import (
	"{{.Path}}/webapp/internal/application"
)

func main() {
	app := application.NewApp()
	app.Start()
}
