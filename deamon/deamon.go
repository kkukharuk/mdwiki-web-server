package deamon

import (
	"fmt"
	"github.com/mister87/mdwiki-web-server/ui"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Host string
	Port int
}

func (cfg *Config) Run() error {
	listenSpec := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	l, err := net.Listen("tcp", listenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}
	log.Printf("Starting, HTTP on: %s\n", l.Addr().String())
	app := ui.Start(l)
	waitForSignal(app)
	return nil
}

func waitForSignal(app *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.\n", s)
	ui.Stop(app)
}
