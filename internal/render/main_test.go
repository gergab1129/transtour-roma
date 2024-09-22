package render

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

type templateFile struct {
	name     string
	contents string
}

type testWriter struct{}

func (tw testWriter) Header() http.Header {
	return http.Header{}
}

func (tw testWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (tw testWriter) WriteHeader(statusCode int) {
}

var Rend *Renderer

func TestMain(m *testing.M) {
	files := []templateFile{
		{"base.layout.tmpl", `{{ define "base" }} This is the base layout {{ block "content" . }} {{ end }} end of base layout {{ end }}\n`},
		{"content.pages.tmpl", `{{ template "base" }} {{ define "content" }} here goes the content {{ end }}\n`},
		{"about.pages.tmpl", `{{ template "base" }} {{define "content"}} this is about us {{ end }}\n`},
		{"other.pages.tmpl", `{{ template "base" }} this is about us\n`},
	}

	createTestDir("../../static/templates/", files)
	Rend = New()
	os.Exit(m.Run())
}

func createTestDir(templateDir string, files []templateFile) {

	for _, file := range files {
		f, err := os.Create(filepath.Join(templateDir, file.name))
		if err != nil {
			log.Fatalf("failed creating testing file %s. error=%s", file.name, err)
		}

		defer f.Close()
		_, err = io.WriteString(f, file.contents)

		if err != nil {
			log.Fatalf("failed writing contents to testing file %s. error=%s", file.name, err)
		}
	}

}
