package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/niteshchandra7/bookings/internals/models"
)

// AppConfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
