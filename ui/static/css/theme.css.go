package css

import (
	"encoding/base64"
	"net/http"
)

func (cfg *Config) ThemeCSS(w http.ResponseWriter, r *http.Request) {
	cfg.Logger.Debugf("Getting page: %s", r.URL.Path)
	css_b64 := "Ym9keSB7CiAgcGFkZGluZy10b3A6IDcwcHg7CiAgcGFkZGluZy1ib3R0b206IDMwcHg7Cn0KCi50aGVtZS1kcm9wZG93biAuZHJvcGRvd24tbWVudSB7CiAgcG9zaXRpb246IHN0YXRpYzsKICBkaXNwbGF5OiBibG9jazsKICBtYXJnaW4tYm90dG9tOiAyMHB4Owp9CgoudGhlbWUtc2hvd2Nhc2UgPiBwID4gLmJ0biB7CiAgbWFyZ2luOiA1cHggMDsKfQoKLnRoZW1lLXNob3djYXNlIC5uYXZiYXIgLmNvbnRhaW5lciB7CiAgd2lkdGg6IGF1dG87Cn0K"
	css, _ := base64.StdEncoding.DecodeString(css_b64)
	w.Header().Set("Content-Type", "text/css")
	w.Write(css)
}
