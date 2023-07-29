package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/niteshchandra7/bookings/internals/config"
	"github.com/niteshchandra7/bookings/internals/handlers"
	"github.com/niteshchandra7/bookings/internals/helpers"
	"github.com/niteshchandra7/bookings/internals/models"
	"github.com/niteshchandra7/bookings/internals/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	run()
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application of port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("error while starting the server")
		os.Exit(1)
	}
}

func run() error {

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplate()

	if err != nil {
		log.Fatal("cannot create template cache", err)
		return err
	}
	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
