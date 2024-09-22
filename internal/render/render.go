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
	templatePath = "templates/"
	layoutFile   = "templates/base.layout.tmpl"
)

var appConfig *config.AppConfig

type Renderer struct {
	Render        func(out http.ResponseWriter, tmpl string, tc map[string]*template.Template) error
	TemplateCache map[string]*template.Template
}

func New(conf *config.AppConfig) *Renderer {
	tc, err := templateCache()
	if err != nil {
		log.Fatalf("failed to create template cache. got=%s", err)
	}
	appConfig = conf

	return &Renderer{
		Render:        render,
		TemplateCache: tc,
	}
}

func DefaultData() {
}

func render(out http.ResponseWriter, tmpl string, tc map[string]*template.Template) error {
	if !appConfig.InProduction {
		var err error
		tc, err = templateCache()
		if err != nil {
			return err
		}
	}

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

func templateCache() (map[string]*template.Template, error) {
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
