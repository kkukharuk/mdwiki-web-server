package ui

import (
	"github.com/mister87/mdwiki-web-server/logger"
	"github.com/mister87/mdwiki-web-server/ui/handlers"
	"github.com/mister87/mdwiki-web-server/ui/static/css"
	"github.com/mister87/mdwiki-web-server/ui/static/js"
	"net"
	"net/http"
	"time"
)

type Config struct {
	Listener net.Listener
	Logger   logger.Config
	UI       *http.Server
}

func (cfg *Config) Start(mdPath string) {
	cfg.Logger.Debugf("Configuring HTTP server on: %s", cfg.Listener.Addr().String())
	server := &http.Server{
		Addr:           cfg.Listener.Addr().String(),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	cfg.UI = server
	cfg.Logger.Debug("Prepare HTTP server handlers")
	handlersCfg := handlers.Config{
		MDWikiPath: mdPath,
		Logger:     cfg.Logger,
	}
	// Pages
	http.Handle("/", handlersCfg.Other())
	http.Handle("/ui", handlersCfg.Index())
	http.Handle("/login", handlersCfg.Login())
	// Static Handles
	cssConfig := css.Config{
		Logger: cfg.Logger,
	}
	http.HandleFunc("/static/css/bootstrap.min.css", cssConfig.BootstrapMinCSS)
	http.HandleFunc("/static/css/ie10-viewport-bug-workaround.css", cssConfig.IE10ViewportBugWorkaroundCSS)
	http.HandleFunc("/static/css/signin.css", cssConfig.SigninCSS)
	http.HandleFunc("/static/css/theme.css", cssConfig.ThemeCSS)
	jsConfig := js.Config{
		Logger: cfg.Logger,
	}
	http.HandleFunc("/static/js/html5shiv.min.js", jsConfig.Html5shivMinJS)
	http.HandleFunc("/static/js/ie8-responsive-file-warning.js", jsConfig.IE8ResponsiveFileWarningJS)
	http.HandleFunc("/static/js/ie10-viewport-bug-workaround.js", jsConfig.IE10ViewportBugWorkaroundJS)
	http.HandleFunc("/static/js/ie-emulation-modes-warning.js", jsConfig.IEEmulationModesWarningJS)
	http.HandleFunc("/static/js/respond.min.js", jsConfig.RespondMinJS)
	cfg.Logger.Debugf("Starting, HTTP server on: %s", cfg.Listener.Addr().String())
	go server.Serve(cfg.Listener)
}

func (cfg *Config) Stop() error {
	cfg.Logger.Debugf("Stopping, HTTP server on: %s", cfg.UI.Addr)
	err := cfg.UI.Shutdown(nil)
	if err != nil {
		cfg.Logger.Errorf("Error stoping HTTP server: %v", err)
	}
	return err
}
