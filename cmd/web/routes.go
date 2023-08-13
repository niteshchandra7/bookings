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
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(Auth)
	//mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)

	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	mux.Get("/book-room", handlers.Repo.BookRoom)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.AdminDashboard)
	})
	return mux
}
