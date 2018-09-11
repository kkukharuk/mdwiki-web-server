package deamon

import (
	"fmt"
	"github.com/mister87/mdwiki-web-server/logger"
	"github.com/mister87/mdwiki-web-server/ui"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Host   string
	Port   int
	Logger logger.Config
}

func (cfg *Config) Run(mdPath string) error {
	//cfg.createPid()
	cfg.Logger.Infof("Create listener: %s:%d", cfg.Host, cfg.Port)
	listenSpec := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	l, err := net.Listen("tcp", listenSpec)
	if err != nil {
		cfg.Logger.Errorf("Error creating listener: %v", err)
		return err
	}
	app := ui.Config{
		Listener: l,
		Logger:   cfg.Logger,
	}
	app.Start(mdPath)
	return cfg.waitForSignal(app)
}

func (cfg *Config) waitForSignal(app ui.Config) error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	cfg.Logger.Infof("Got signal: %v, exiting.", s)
	return app.Stop()
}

func (cfg *Config) createPid() {
	pid, err := os.OpenFile("mdwiki-web-server.pid", os.O_WRONLY|os.O_CREATE, 0600)
	defer pid.Close()
	if err != nil {
		cfg.Logger.Errorf("Error create pid-file: %v", err)
	}
	_, err = pid.WriteString(fmt.Sprintf("%d", os.Getgid()))
	if err != nil {
		cfg.Logger.Errorf("Error write pid number in pid-file: %v", err)
	}
}
