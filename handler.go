package coco

import (
	"encoding/json"
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
	w             http.ResponseWriter
	headerWritten bool

	ctx Context
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
	r.headerWritten = true
}

// JSON returns a JSON response
func (r *Response) JSON(statusCode int, v interface{}) error {
	r.w.Header().Set("Content-Type", "application/json")

	if !r.headerWritten {
		r.WriteHeader(statusCode)
	}
	return json.NewEncoder(r.w).Encode(v)
}

// Text returns a text response
func (r *Response) Text(statusCode int, v string) error {
	r.w.Header().Set("Content-Type", "text/plain")

	if !r.headerWritten {
		r.WriteHeader(statusCode)
	}
	_, err := r.w.Write([]byte(v))
	return err
}

// Render returns a HTML response with a template
func (r *Response) Render(name string, data interface{}) error {

	store := r.ctx.templateStore
	if store == nil {
		return fmt.Errorf("no templates loaded")
	}

	fmt.Println("rendering template: ", name)

	tmpl, err := store.Get(name)
	if err != nil {
		return err
	}

	fmt.Println(tmpl.Name())

	return tmpl.Execute(r.w, data)

}
