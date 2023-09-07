package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vikash-parashar/bookings/pkg/config"
	"github.com/vikash-parashar/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	// prevent our app from cross site forgery
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/major-suit", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// setting up file server to serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
