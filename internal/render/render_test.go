package render

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

type templateFile struct {
	name     string
	contents string
}

func Test_TemplateCache(t *testing.T) {
	files := []templateFile{
		{"base.layout.tmpl", `{{ define "base" }} This is the base layout {{ block "content" . }} {{ end }} end of base layout {{ end }}\n`},
		{"content.pages.tmpl", `{{ template "base" }} {{ define "content" }} here goes the content {{ end }}\n`},
		{"about.pages.tmpl", `{{ template "base" }} {{define "content"}} this is about us {{ end }}\n`},
		{"other.pages.tmpl", `{{ template "base" }} this is about us\n`},
	}

	createTestDir(t, "../../static/templates/", files)

	tc, err := TemplateCache()
	if err != nil {
		t.Errorf("template cache creation failed. got=%s", err)
		t.FailNow()
	}

	for _, tt := range files {
		tmpl, ok := tc[tt.name]

		if !ok && (tt.name != "base.layout.tmpl") {
			t.Errorf("%s not stored in template cache", tt.name)
		}

		if tt.name != "base.layout.tmpl" {
			err = tmpl.Execute(os.Stderr, nil)
			if err != nil {
				t.Errorf("template %s execution failed. got=%s", tt.name, err)
			}
		}

	}
}

func createTestDir(t *testing.T, templateDir string, files []templateFile) {

	for _, file := range files {
		f, err := os.Create(filepath.Join(templateDir, file.name))
		if err != nil {
			t.Errorf("failed creating testing file %s. error=%s", file.name, err)
			t.FailNow()
		}

		defer f.Close()
		_, err = io.WriteString(f, file.contents)

		if err != nil {
			t.Errorf("failed writing contents to testing file %s. error=%s", file.name, err)
			t.FailNow()
		}
	}

}
