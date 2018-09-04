package handlers

import (
	"fmt"
	"net/http"
)

func (cfg *Config) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	cfg.Logger.Warnf("Page '%s' not found", r.URL.Path)
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
