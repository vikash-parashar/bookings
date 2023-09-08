package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikash-parashar/bookings/internal/config"
	"github.com/vikash-parashar/bookings/internal/handlers"
	"github.com/vikash-parashar/bookings/internal/models"
	"github.com/vikash-parashar/bookings/internal/render"
)

const (
	port = ":8080"
)

var sessionManager *scs.SessionManager

var app config.AppConfig

func main() {
	err := Run()
	if err != nil {
		log.Fatalln(err)
	}
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start application")
		return
	}
}

var InfoLog *log.Logger
var ErrorLog *log.Logger

func Run() error {
	// what i am going to put in the sessions
	gob.Register(models.Reservation{})

	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = InfoLog

	ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = ErrorLog

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
		return err
	}
	app.TemplateCache = tc
	app.UseCache = true // keep false if you are in developer mod
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	log.Println("starting our application on port ", port)
	return nil
}
