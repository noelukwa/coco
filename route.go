package coco

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	base string
	hr   *httprouter.Router

	// Middleware
	middleware []Handler
}

func (r *Route) Use(middleware ...Handler) *Route {
	r.middleware = append(r.middleware, middleware...)
	return r
}

func (r *Route) combineHandlers(handlers ...Handler) []Handler {
	return append(r.middleware, handlers...)
}

func (a *App) NewRoute(path string) *Route {
	if r, ok := a.routes[path]; ok {
		return &r
	}

	return &Route{
		base: path,
		hr:   a.base,
	}
}

func (r *Route) handle(httpMethod string, path string, handlers []Handler) {
	handlers = r.combineHandlers(handlers...)
	r.hr.Handle(httpMethod, r.base+path, func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		c := Context{
			handlers: handlers,
		}

		response := Response{w}
		request := Request{req, p}

		c.next(response, &request)
	})
}

// GET method
func (r *Route) Get(path string, handlers ...Handler) {
	r.handle("GET", path, handlers)
}
