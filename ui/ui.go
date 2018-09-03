package ui

import (
	"github.com/mister87/mdwiki-web-server/ui/handlers"
	"github.com/mister87/mdwiki-web-server/ui/static/css"
	"github.com/mister87/mdwiki-web-server/ui/static/js"
	"log"
	"net"
	"net/http"
	"time"
)

func Start(listener net.Listener) *http.Server {
	server := &http.Server{
		Addr:           listener.Addr().String(),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	// Pages
	http.Handle("/", handlers.Other())
	http.Handle("/ui", handlers.Index())
	http.Handle("/login", handlers.Login())
	// Static Handles
	http.HandleFunc("/static/css/bootstrap.min.css", css.BootstrapMinCSS)
	http.HandleFunc("/static/css/ie10-viewport-bug-workaround.css", css.IE10ViewportBugWorkaroundCSS)
	http.HandleFunc("/static/css/signin.css", css.SigninCSS)
	http.HandleFunc("/static/css/theme.css", css.ThemeCSS)
	http.HandleFunc("/static/js/html5shiv.min.js", js.Html5shivMinJS)
	http.HandleFunc("/static/js/ie8-responsive-file-warning.js", js.IE8ResponsiveFileWarningJS)
	http.HandleFunc("/static/js/ie10-viewport-bug-workaround.js", js.IE10ViewportBugWorkaroundJS)
	http.HandleFunc("/static/js/ie-emulation-modes-warning.js", js.IEEmulationModesWarningJS)
	http.HandleFunc("/static/js/respond.min.js", js.RespondMinJS)
	go server.Serve(listener)
	return server
}

func Stop(app *http.Server) error {
	err := app.Shutdown(nil)
	log.Printf("Stopping, HTTP on: %s\n", app.Addr)
	return err
}
