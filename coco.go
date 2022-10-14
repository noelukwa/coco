package coco

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// NextFunc is a function that is called to pass execution to the next handler
type NextFunc func(rw Response, r *Request)

// Handle is a wrapper for httprouter.Handle
type Handler func(rw Response, req *Request, next NextFunc)

// TODO: support template rendering
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

	//TODO: template store
	templateStore templateStore
}

// GlobalPrefix sets a global prefix for all routes
func (a *App) GlobalPrefix(prefix string) {
	a.basePath = strings.TrimSpace(strings.TrimFunc(prefix, func(r rune) bool { return r == '/' }))
}

func NewApp() (app *App) {
	router := httprouter.New()
	app = &App{
		basePath: "/",
		routes:   make(map[string]*Route),
		base:     router,
	}
	app.Route = app.NewRoute(app.basePath)
	return
}

//TODO: LoadTemplates
func (app *App) LoadTemplates(config TemplateConfig) error {

	err := fs.WalkDir(config.Path, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

//TODO: ServeStatic serves static files from a given directory
func (a *App) ServeStatic(path string, dir fs.FS) {

	// a.static.Get(path, func(rw Response, r *Request, _ NextFunc) {
	// 	http.ServeFile(rw, r.Request, a.static.pathToFiles)
	// })

}

// Listen starts an HTTP server
func (a *App) Listen(addr string) {
	http.ListenAndServe(addr, a)
}

// ListenTLS starts an HTTPS server
func (a *App) ListenTLS(addr string, certFile string, keyFile string) {
	http.ListenAndServeTLS(addr, certFile, keyFile, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.base.ServeHTTP(w, r)
}
