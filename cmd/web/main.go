package main

import (
	"Booking/internal/config"
	"Booking/internal/handlers"
	"Booking/internal/helpers"
	"Booking/internal/models"
	"Booking/internal/render"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var appConfiguration config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	servingPage := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfiguration),
	}

	err = servingPage.ListenAndServe()
	log.Fatal(err)
}

func run() error {

	//What am i going to true when in production
	gob.Register(models.Reservation{})

	appConfiguration.InProduction = false //this should be true in production

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfiguration.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfiguration.InfoLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfiguration.InProduction

	appConfiguration.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	appConfiguration.TemplateCache = templateCache
	// while developing application make the field value false
	// in production keep the value true
	appConfiguration.UseCache = false

	repo := handlers.NewRepo(&appConfiguration)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfiguration)
	helpers.NewHelpers(&appConfiguration)

	return nil
}
