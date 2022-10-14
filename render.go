package coco

import (
	"html/template"
	"io/fs"
)

type templateStore struct {
	templates map[string]*template.Template

	templateFiles map[string]bool

	includeFiles []string

	sharedFiles []string

	ext string

	path fs.FS
}

// TemplateConfig tells coco how to group template files for parsing.
type TemplateConfig struct {
	// The directory where shared templates are stored.
	// Shared templates are templates that are used by more than one page.
	// For example, a layout template.
	// If this is not set, then the default directory is "shared".
	Globals []string

	// The extension of the template files.
	// Defaults to ".html".
	// Only extensions supported by the html/template package are allowed.
	Ext string

	// The location of the template files relative to the directory
	Path fs.FS

	// List of files or directories that contain partial templates.
	// Partial templates are templates that are not rendered directly.
	// For example: navigation, footer, etc.
	// If not specified, a directory named "includes" is used.
	Includes []string
}

// Get returns a template by name if found.
func (t *templateStore) Get(name string) (*template.Template, error) {
	if t.templates == nil {
		t.templates = make(map[string]*template.Template)
	}

	if t.templates[name] != nil {
		return t.templates[name], nil
	}

	nameWithExt := name + t.ext

	// check if the templatefiles

	parseTarget := []string{nameWithExt}
	parseTarget = append(parseTarget, t.includeFiles...)
	parseTarget = append(parseTarget, t.sharedFiles...)

	tmpl, err := template.New(nameWithExt).ParseFS(t.path, parseTarget...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
