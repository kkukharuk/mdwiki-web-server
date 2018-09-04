package handlers

import (
	"fmt"
	"net/http"
)

func (cfg *Config) ErrorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
