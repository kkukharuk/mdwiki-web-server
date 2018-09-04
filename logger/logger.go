package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	LogLevel string
	LogFile  string
}

func (cfg *Config) Init() error {
	f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.LogLevel != "DEBUG" && cfg.LogLevel != "" && cfg.LogLevel != "WARN" && cfg.LogLevel != "INFO" {
		log.Fatalf("Error setting log level: %s", cfg.LogLevel)
	}
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)
	return nil
}

func (cfg *Config) Error(v ...interface{}) {
	msg := "[ERROR] " + fmt.Sprint(v...)
	log.Print(msg)
}

func (cfg *Config) Errorf(format string, v ...interface{}) {
	msg := "[ERROR] " + fmt.Sprintf(format, v...)
	log.Print(msg)
}

func (cfg *Config) Info(v ...interface{}) {
	msg := "[INFO] " + fmt.Sprint(v...)
	log.Print(msg)
}

func (cfg *Config) Infof(format string, v ...interface{}) {
	msg := "[INFO] " + fmt.Sprintf(format, v...)
	log.Print(msg)
}

func (cfg *Config) Warn(v ...interface{}) {
	msg := "[WARN] " + fmt.Sprint(v...)
	log.Print(msg)
}

func (cfg *Config) Warnf(format string, v ...interface{}) {
	msg := "[WARN] " + fmt.Sprintf(format, v...)
	log.Print(msg)
}

func (cfg *Config) Debug(v ...interface{}) {
	msg := "[DEBUG] " + fmt.Sprint(v...)
	log.Print(msg)
}

func (cfg *Config) Debugf(format string, v ...interface{}) {
	msg := "[DEBUG] " + fmt.Sprintf(format, v...)
	log.Print(msg)
}
