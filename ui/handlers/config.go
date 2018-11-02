package handlers

import "github.com/mister87/mdwiki-web-server/logger"

type Config struct {
	MDWikiPath string
	Logger     logger.Config
	Auth       bool
}
