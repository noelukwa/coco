package coco

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

// NextFunc is a function that is called to pass execution to the next handler
type NextFunc func(rw Response, r *Request)

// Handle is a wrapper for httprouter.Handle
type Handler func(rw Response, req *Request, next NextFunc)

type StaticRoute struct {
	*Route
	pathToFiles string
}

type App struct {
	base     *httprouter.Router
	basePath string
	*Route   // default route
	routes   map[string]*Route

	//TODO: static files
	static StaticRoute

	templates map[string]*template.Template
}

// GlobalPrefix sets a global prefix for all routes
func (a *App) GlobalPrefix(prefix string) *App {
	if prefix == "" {
		return a
	}
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}
	a.Route.base = prefix
	a.basePath = path.Clean(prefix)
	return a
}

func NewApp() (app *App) {
	app = &App{
		basePath: "",
		routes:   make(map[string]*Route),
		base:     httprouter.New(),
	}
	app.Route = app.NewRoute(app.basePath)
	return
}

//TODO: ServeStatic serves static files from a given directory
func (a *App) ServeStatic(path string, dir fs.FS) {

	// a.static.Get(path, func(rw Response, r *Request, _ NextFunc) {
	// 	http.ServeFile(rw, r.Request, a.static.pathToFiles)
	// })

}

// Listen starts an HTTP server
func (a *App) Listen(addr string, ctx context.Context) error {

	server := &http.Server{
		Addr:    addr,
		Handler: a,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("shutting down server")
		server.Shutdown(ctx)
	}()

	return server.ListenAndServe()
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.base.ServeHTTP(w, r)
}
