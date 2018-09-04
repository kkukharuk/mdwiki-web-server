package main

import (
	"flag"
	"github.com/mister87/mdwiki-web-server/deamon"
)

var app deamon.Config

func init() {
	flag.StringVar(&app.Host, "h", "", "HTTP host")
	flag.IntVar(&app.Port, "p", 3000, "HTTP port")
	flag.StringVar(&app.MDWikiPath, "wiki", "", "Markdown files path")
	flag.StringVar(&app.Logger.LogFile, "log", "mdwiki-web-server.log", "Log file path")
	flag.StringVar(&app.Logger.LogLevel, "log-level", "DEBUG", "Log level: DEBUG, ERROR, WARN, INFO")
	if app.MDWikiPath == "" {
	}
	app.Logger.Init()
	flag.Parse()
}

func main() {
	if err := app.Run(); err != nil {
		app.Logger.Errorf("Error in main(): %v", err)
	}
}
