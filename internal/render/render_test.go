package render

import (
	"os"
	"testing"
)

var tw testWriter

func Test_render(t *testing.T) {
	tests := []string{
		"content.pages.tmpl",
		"about.pages.tmpl",
		"other.pages.tmpl",
	}

	for _, tt := range tests {

		err := Rend.Render(tw, tt, Rend.TemplateCache)
		if err != nil {
			t.Errorf("error executing render function for test=%s. got error=%s", tt, err)
		}
	}
}

func Test_TemplateCache(t *testing.T) {
	files := []string{
		"content.pages.tmpl",
		"about.pages.tmpl",
		"other.pages.tmpl",
	}

	tc, err := templateCache()
	if err != nil {
		t.Errorf("template cache creation failed. got=%s", err)
		t.FailNow()
	}

	for _, tt := range files {
		tmpl, ok := tc[tt]

		if !ok && (tt != "base.layout.tmpl") {
			t.Errorf("%s not stored in template cache", tt)
		}

		if tt != "base.layout.tmpl" {
			err = tmpl.Execute(os.Stderr, nil)
			if err != nil {
				t.Errorf("template %s execution failed. got=%s", tt, err)
			}
		}

	}
}
