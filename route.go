package coco

import (
	"net/http"
	"path"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	base string
	hr   *httprouter.Router

	// Middleware
	middleware []Handler

	app *App
}

func (r *Route) Use(middleware ...Handler) *Route {
	r.middleware = append(r.middleware, middleware...)
	return r
}

func (r *Route) combineHandlers(handlers ...Handler) []Handler {
	return append(r.middleware, handlers...)
}

// pathify func turns whatever string to path just to be sure.
func (a *App) pathify(p string) string {
	return a.basePath + path.Clean(p)
}

func (a *App) NewRoute(path string) *Route {
	if path == "" {
		path = "/"
	}
	path = a.pathify(path)
	if r, ok := a.routes[path]; ok {
		return r
	}
	r := &Route{
		base: path,
		hr:   a.base,
		app:  a,
	}
	a.routes[path] = r

	return r
}

// getfullPath method returns the full path of the route
func (r *Route) getfullPath(path string) string {

	raw := strings.Trim(path, "/")

	if len(raw) > 0 && raw[0] == ':' {
		return r.base + "/" + raw
	}

	return r.base + strings.TrimPrefix(path, "/")
}

func (r *Route) handle(httpMethod string, path string, handlers []Handler) {
	handlers = r.combineHandlers(handlers...)
	r.hr.Handle(httpMethod, r.getfullPath(path), func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		c := Context{
			handlers:  handlers,
			templates: r.app.templates,
		}
		response := Response{w, false, c}
		request := Request{req, p}

		c.next(response, &request)
	})
}

// GET method
func (r *Route) Get(path string, handlers ...Handler) {
	r.handle("GET", path, handlers)
}

// POST method
func (r *Route) Post(path string, handlers ...Handler) {
	r.handle("POST", path, handlers)
}

// PUT method
func (r *Route) Put(path string, handlers ...Handler) {
	r.handle("PUT", path, handlers)
}

// DELETE method
func (r *Route) Delete(path string, handlers ...Handler) {
	r.handle("DELETE", path, handlers)
}

// PATCH method
func (r *Route) Patch(path string, handlers ...Handler) {
	r.handle("PATCH", path, handlers)
}

// OPTIONS method
func (r *Route) Options(path string, handlers ...Handler) {
	r.handle("OPTIONS", path, handlers)
}

// HEAD method
func (r *Route) Head(path string, handlers ...Handler) {
	r.handle("HEAD", path, handlers)
}

// Any method
func (r *Route) Any(path string, handlers ...Handler) {
	r.handle("GET", path, handlers)
	r.handle("POST", path, handlers)
	r.handle("PUT", path, handlers)
	r.handle("DELETE", path, handlers)
	r.handle("PATCH", path, handlers)
	r.handle("OPTIONS", path, handlers)
	r.handle("HEAD", path, handlers)
}

// Group method
func (r *Route) Group(path string, handlers ...Handler) *Route {
	return &Route{
		base:       r.base + path,
		hr:         r.hr,
		middleware: r.combineHandlers(handlers...),
	}
}

// Static method
func (r *Route) Static(path string, root string) {

	r.hr.ServeFiles(r.base+path, http.Dir(root))
}

// File method
func (r *Route) File(path string, file string) {

	r.hr.ServeFiles(r.base+path, http.Dir(file))
}
