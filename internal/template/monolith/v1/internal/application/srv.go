package application

import (
	"github.com/golangee/forms/theme/material"
	gzipped "github.com/lpar/gzipped/v2"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"{{.Path}}/internal/build"
)


func withIndexHTML(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			newpath := path.Join(r.URL.Path, "index.html")
			r.URL.Path = newpath
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) startSrv(dir string, port int) {
	log.Println("build time", build.Time)
	log.Println("build commit", build.Commit)

	material.Resources(http.DefaultServeMux)

//	srv := http2.NewServer()
//	srv.SetNotFound(withIndexHTML(gzipped.FileServer(gzipped.Dir(dir))))
//	a.initControllers(srv)
//	http.Handle("/", srv.Handler())
	http.Handle("/",withIndexHTML(gzipped.FileServer(gzipped.Dir(dir))))

	log.Printf("Serving %s on HTTP port: %d\n", dir, port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
