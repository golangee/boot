package application

import (
	. "github.com/golangee/forms"
	"github.com/golangee/log"
	"{{.Path}}/webapp/internal/build"
	"{{.Path}}/webapp/internal/notfound"
)

type App struct {
	*Application
}


func NewApp() *App {
	a := &App{}
	logger := log.New("")
	logger.Info("{{.BinaryName}} frontend", log.Obj("version", build.Env().String()))
	a.Application = NewApplication(a, build.Env().String())
	return a
}

func (a *App) Start() {
	Theme().SetColor(0x1b8c30ff)

	a.UnmatchedRoute(notfound.FromQuery)
	a.Application.Start()
}
