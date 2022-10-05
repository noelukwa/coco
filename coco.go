package coco

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NextFunc func(rw Response, r *Request)

type Handler func(rw Response, req *Request, next NextFunc)

type App struct {
	base   *httprouter.Router
	*Route // default route
	routes map[string]Route
}

func NewApp() *App {
	router := httprouter.New()
	routes := make(map[string]Route)
	routes["/"] = Route{base: "/", hr: router}
	defaultRoute := routes["/"]
	return &App{
		base:   router,
		routes: routes,
		Route:  &defaultRoute,
	}
}

func (a *App) Listen(addr string) {
	http.ListenAndServe(addr, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.base.ServeHTTP(w, r)
}
