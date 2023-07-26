package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/niteshchandra7/bookings/internals/config"
	"github.com/niteshchandra7/bookings/internals/handlers"
	//"github.com/go-chi/chi/"
	//"github.com/go-chi/chi/middleware"
	//"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(WriteToConsole)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Get("/generals-quarters", http.HandlerFunc(handlers.Repo.Generals))
	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))
	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Get("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
