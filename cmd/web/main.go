package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/niteshchandra7/bookings/internals/config"
	"github.com/niteshchandra7/bookings/internals/drivers"
	"github.com/niteshchandra7/bookings/internals/handlers"
	"github.com/niteshchandra7/bookings/internals/helpers"
	"github.com/niteshchandra7/bookings/internals/models"
	"github.com/niteshchandra7/bookings/internals/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application of port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error while starting the server")
		os.Exit(1)
	}
}

func run() (*drivers.DB, error) {
	// change this true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// what I am going to put in session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	log.Println("Connecting to database...")

	db, err := drivers.ConnectSQL("host=localhost port=5432 dbname=bookings user=niteshchandra password=")
	if err != nil {
		log.Fatal("cannot connect to db, dying...")
		return nil, err
	}

	tc, err := render.CreateTemplate()

	if err != nil {
		log.Fatal("cannot create template cache", err)
		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db, nil
}
