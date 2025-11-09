package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KiroLakestrike/bedAndBreakfast/pkg/config"
	"github.com/KiroLakestrike/bedAndBreakfast/pkg/handlers"
	"github.com/KiroLakestrike/bedAndBreakfast/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true when in PRoduction
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.PortNumber = ":8080"

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Load Handlers
	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes(&app),
	}

	fmt.Println(fmt.Sprintf("Listening on http://localhost%v", app.PortNumber))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
