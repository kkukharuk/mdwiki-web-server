package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Config struct {
	mu           sync.Mutex
	LogLevel     string
	LogFileName  string
	LogFile      *os.File
	Log          *log.Logger
	RotateConfig RotateRule
	Output       bool
}

func (cfg *Config) Init() error {
	if cfg.LogLevel != "DEBUG" && cfg.LogLevel != "" && cfg.LogLevel != "WARN" && cfg.LogLevel != "INFO" {
		log.Fatalf("[ERROR] Error setting log level: %s", cfg.LogLevel)
	}
	if cfg.RotateConfig.RuleType != "SIZE" && cfg.RotateConfig.RuleType != "DATE" {
		log.Fatalf("[ERROR] Error setting log rotate type: %s", cfg.RotateConfig.RuleType)
	}
	cfg.RotateConfig.Increment = 0
	cfg.Log = log.New(cfg.GetWriter(), "", log.Ldate|log.Ltime)
	return nil
}

func (cfg *Config) GetWriter() io.Writer {
	var err error
	cfg.LogFile, err = os.OpenFile(cfg.LogFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	if cfg.Output {
		return io.MultiWriter(cfg.LogFile, os.Stdout)
	}
	return cfg.LogFile
}

func (cfg *Config) Error(v ...interface{}) {
	cfg.Rotate()
	msg := "[ERROR] " + fmt.Sprint(v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Errorf(format string, v ...interface{}) {
	cfg.Rotate()
	msg := "[ERROR] " + fmt.Sprintf(format, v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Info(v ...interface{}) {
	cfg.Rotate()
	msg := "[INFO] " + fmt.Sprint(v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Infof(format string, v ...interface{}) {
	cfg.Rotate()
	msg := "[INFO] " + fmt.Sprintf(format, v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Warn(v ...interface{}) {
	cfg.Rotate()
	msg := "[WARN] " + fmt.Sprint(v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Warnf(format string, v ...interface{}) {
	cfg.Rotate()
	msg := "[WARN] " + fmt.Sprintf(format, v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Debug(v ...interface{}) {
	cfg.Rotate()
	msg := "[DEBUG] " + fmt.Sprint(v...)
	cfg.Log.Print(msg)
}

func (cfg *Config) Debugf(format string, v ...interface{}) {
	cfg.Rotate()
	msg := "[DEBUG] " + fmt.Sprintf(format, v...)
	cfg.Log.Print(msg)
}
