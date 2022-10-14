package coco

import "net/http"

// TODO: add support for file upload
type File struct {
	Filename string
	Header   http.Header
	Size     int64
}

type FileUpload struct {
	FieldName string
	MaxSize   int64
	Allowed   []string
	Store     string
	SaveFunc  func(file *File) string
}

// Single method tells upload handler to only accept a single file
func (f *FileUpload) Single(rw Response, r *Request, next NextFunc) {
	f.MaxSize = 1
	next(rw, r)
}

// Multi method tells upload handler to accept multiple files
func (f *FileUpload) Multi(rw Response, r *Request, next NextFunc) {
	f.MaxSize = 0
	next(rw, r)
}
