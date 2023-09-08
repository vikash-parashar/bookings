package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikash-parashar/bookings/internal/config"
	"github.com/vikash-parashar/bookings/internal/handlers"
	"github.com/vikash-parashar/bookings/internal/render"
)

const (
	port = ":8080"
)

var sessionManager *scs.SessionManager

var app config.AppConfig

func main() {

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.InProduction = false // keep false if you are in developer mod
	app.SessionManager = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = true // keep false if you are in developer mod
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	log.Println("starting our application on port ", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start application")
		return
	}
}
