package config

import (
	"database/sql"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	InfoLog        *log.Logger
	Db             *sql.DB
	InProduction   bool
	SessionManager *scs.SessionManager
}
