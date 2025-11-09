package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// This is just a basic configuation file
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	PortNumber    string
	InProduction  bool
	Session       *scs.SessionManager
}
