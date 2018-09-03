package main

import (
	"flag"
	"github.com/mister87/mdwiki-web-server/deamon"
	"log"
)

var app deamon.Config

func init() {
	flag.StringVar(&app.Host, "h", "", "HTTP host")
	flag.IntVar(&app.Port, "p", 3000, "HTTP port")
	flag.Parse()
}

func main() {
	if err := app.Run(); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
