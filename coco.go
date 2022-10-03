package coco

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NextFunc func(rw Response, r *Request)

// implement ServeHTTP for NextFunc
func (n NextFunc) ServeHTTP(w Response, r *Request) {
	n(w, r)
}

type Handler func(rw Response, req *Request, next NextFunc)

type App struct {
	baseRouter *httprouter.Router
	routes     map[string]Route
}

// default route method for app
func (a *App) Route(path string) *Route {
	if r, ok := a.routes[path]; ok {
		return &r
	}
	return nil
}

func (a *App) defaultRoute() *Route {
	if r, ok := a.routes["/"]; ok {
		return &r
	}
	panic("default route not found")
}

func (a *App) Get(path string, handler Handler) {

	r := a.defaultRoute()
	rp, h := r.getHandle(path, handler)
	r.hr.Handle("GET", rp, h)
}

func NewApp() *App {
	router := httprouter.New()
	routes := make(map[string]Route)
	routes["/"] = Route{base: "/", hr: router}
	return &App{
		baseRouter: router,
		routes:     routes,
	}
}

func (a *App) Listen(addr string) {
	http.ListenAndServe(addr, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.baseRouter.ServeHTTP(w, r)
}
