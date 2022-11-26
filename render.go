package coco

import (
	"fmt"
	"html/template"
	"io/fs"

	"github.com/noelukwa/tempest"
)

type TemplateConfig struct {
	// The file extension of the templates.
	// Defaults to ".html".
	Ext string

	// The directory where the includes are stored.
	// Defaults to "includes".
	IncludesDir string

	// The name used for layout templates :- templates that wrap other contents.
	// Defaults to "layouts".
	Layout string
}

// LoadTemplates loads templates from an fs.FS with a given config
func (app *App) LoadTemplates(fs fs.FS, config *TemplateConfig) (err error) {

	if app.templates == nil {
		app.templates = make(map[string]*template.Template)
	}

	if config != nil {
		app.templates, err = tempest.WithConfig(&tempest.Config{
			Layout:      config.Layout,
			IncludesDir: config.IncludesDir,
			Ext:         config.Ext,
		}).LoadFS(fs)

		return
	} else {
		app.templates, err = tempest.New().LoadFS(fs)

		return
	}
}

func (r *Response) Render(name string, data interface{}) error {

	temp := r.ctx.templates[name]
	if temp == nil {
		return fmt.Errorf("template %s not found", name)
	}

	return temp.Execute(r, data)
}
