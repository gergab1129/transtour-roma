package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"transtour-roma/internal/config"
)

const (
	templatePath = "../../static/templates"
	layoutFile   = "../../static/templates/base.layout.tmpl"
)

type Renderer struct {
	Render        func(out http.ResponseWriter, tmpl string, tc map[string]*template.Template) error
	TemplateCache map[string]*template.Template
}

func New(app *config.AppConfig) *Renderer {
	tc, err := TemplateCache()
	if err != nil {
		log.Fatalf("failed to create template cache. got=%s", err)
	}
	return &Renderer{
		Render:        Render,
		TemplateCache: tc,
	}
}

func DefaultData() {
}

func Render(out http.ResponseWriter, tmpl string, tc map[string]*template.Template) error {
	t, ok := tc[tmpl]
	if !ok {
		return fmt.Errorf("template %s not stored in template cache", tmpl)
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		return fmt.Errorf("error executing template %s got=%s", tmpl, err)
	}

	_, err = buf.WriteTo(out)
	if err != nil {
		return fmt.Errorf(
			"error writing buffered template %s to response writer. error=%s",
			tmpl,
			err,
		)
	}
	return nil
}

func TemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)
	// list *.pages.tmpl files
	templateFiles, err := filepath.Glob(fmt.Sprintf("%s/*.pages.tmpl", templatePath))
	if err != nil {
		return nil, fmt.Errorf("failed listing template page files. error=%s", err)
	}

	for _, fn := range templateFiles {
		tn := filepath.Base(fn)
		tmpl, err := template.New(tn).ParseFiles(layoutFile, fn)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template file %s error=%s", fn, err)
		}

		templateCache[tn] = tmpl
	}

	return templateCache, nil
}
