package coco

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	r      *http.Request
	params httprouter.Params
}

func (r *Request) Header() http.Header {
	return r.r.Header
}

func (r *Request) Method() string {
	return r.r.Method
}

func (r *Request) URL() string {
	return r.r.URL.String()
}

func (r *Request) Params() map[string]string {

	fmt.Printf("params: %v\n", r.params)
	m := make(map[string]string)
	for _, p := range r.params {
		m[p.Key] = p.Value
	}
	return m
}

type Response struct {
	w http.ResponseWriter
}

// implement http.ResponseWriter interface for Response
func (r *Response) Header() http.Header {
	return r.w.Header()
}

func (r *Response) Write(b []byte) (int, error) {
	return r.w.Write(b)
}

func (r *Response) WriteHeader(statusCode int) {
	r.w.WriteHeader(statusCode)
}

type Handle struct {
	method string
	path   string
	next   NextFunc
}
