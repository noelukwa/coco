package coco

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	base string
	hr   *httprouter.Router
	// handles map[string]Handle
}

func (a *App) NewRoute(path string) *Route {
	if r, ok := a.routes[path]; ok {
		return &r
	}

	return &Route{
		base: path,
		hr:   a.baseRouter,
		// handles: make(map[string]Handle),
	}
}

func (r *Route) Get(path string, handler Handler) {
	rp, h := r.getHandle(path, handler)
	r.hr.Handle("GET", rp, h)
}

func (r *Route) getHandle(path string, handler Handler) (fullpath string, handle httprouter.Handle) {

	path = strings.TrimPrefix(path, "/")

	if strings.HasPrefix(path, ":") {
		fullpath = fmt.Sprintf("%s/%s", r.base, path)
	} else {
		fullpath = r.base + path
	}

	fmt.Printf("fullpath: %s\n", fullpath)
	handle = func(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
		fmt.Printf("handle: %s\n", fullpath)
		rw := Response{w: w}
		req := Request{r: rq, params: ps}
		next := NextFunc(func(rw Response, r *Request) {
			fmt.Println("next")
		})
		handler(rw, &req, next)
	}
	return

}
