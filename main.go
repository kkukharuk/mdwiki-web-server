package main

import (
	"flag"
	"github.com/mister87/mdwiki-web-server/deamon"
	"log"
)

var app deamon.Config
var logSize int

func init() {
	flag.StringVar(&app.Host, "h", "", "HTTP host")
	flag.IntVar(&app.Port, "p", 3000, "HTTP port")
	flag.StringVar(&app.MDWikiPath, "wiki", "", "Markdown files path")
	flag.StringVar(&app.Logger.LogFileName, "log", "mdwiki-web-server.log", "Log file path")
	flag.StringVar(&app.Logger.LogLevel, "log-level", "DEBUG", "Log level: DEBUG, ERROR, WARN, INFO")
	flag.StringVar(&app.Logger.RotateConfig.RuleType, "log-rotate-type", "SIZE", "Log rotate type: SIZE, DATE")
	flag.IntVar(&logSize, "log-size", 1048576, "Log size (Used if log rotate type is 'SIZE')")
	flag.IntVar(&app.Logger.RotateConfig.Day, "log-date", 1, "Log date (Used if log rotate type is 'DATE')")
	flag.IntVar(&app.Logger.RotateConfig.MaxFiles, "log-max-files", 5, "Log max log files")
	flag.BoolVar(&app.Logger.Output, "log-output", false, "Output log data")
	flag.Parse()
	app.Logger.RotateConfig.LogSize = int64(logSize)
	if app.MDWikiPath == "" {
	}
	err := app.Logger.Init()
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}

func main() {
	if err := app.Run(); err != nil {
		app.Logger.Errorf("Error in main(): %v", err)
	}
}
