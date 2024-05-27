package config

import "html/template"

type AppConfig struct {
	UseCache      bool
	InProduction  bool
	TemplateCache map[string]*template.Template
}
