package logger

import (
	"fmt"
	"log"
	"os"
)

type RotateRule struct {
	RuleType  string
	LogSize   int64
	Day       int
	Increment int
	MaxFiles  int
}

func (cfg *Config) Rotate() error {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	f, err := cfg.LogFile.Stat()
	if err != nil {
		cfg.Errorf("Error check log file: %v", err)
		return err
	}
	if cfg.RotateConfig.RuleType == "SIZE" {
		if f.Size() >= cfg.RotateConfig.LogSize {
			err = cfg.LogFile.Close()
			if err != nil {
				cfg.Errorf("Error rotate log file: %v", err)
				return err
			}
			cfg.RotateConfig.Increment++
			err = os.Rename(cfg.LogFileName, fmt.Sprintf("%s.%d", cfg.LogFileName, cfg.RotateConfig.Increment))
			if err != nil {
				log.Fatalf("[ERROR] %v", err)
			}
			cfg.Log = log.New(cfg.GetWriter(), "", log.Ldate|log.Ltime)
		}
	}
	return nil
}
