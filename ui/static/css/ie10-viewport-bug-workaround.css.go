package css

import (
	"encoding/base64"
	"log"
	"net/http"
)

func IE10ViewportBugWorkaroundCSS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting page: %s", r.URL.Path)
	css_b64 := "LyohCiAqIElFMTAgdmlld3BvcnQgaGFjayBmb3IgU3VyZmFjZS9kZXNrdG9wIFdpbmRvd3MgOCBidWcKICogQ29weXJpZ2h0IDIwMTQtMjAxNSBUd2l0dGVyLCBJbmMuCiAqIExpY2Vuc2VkIHVuZGVyIE1JVCAoaHR0cHM6Ly9naXRodWIuY29tL3R3YnMvYm9vdHN0cmFwL2Jsb2IvbWFzdGVyL0xJQ0VOU0UpCiAqLwoKLyoKICogU2VlIHRoZSBHZXR0aW5nIFN0YXJ0ZWQgZG9jcyBmb3IgbW9yZSBpbmZvcm1hdGlvbjoKICogaHR0cDovL2dldGJvb3RzdHJhcC5jb20vZ2V0dGluZy1zdGFydGVkLyNzdXBwb3J0LWllMTAtd2lkdGgKICovCkAtbXMtdmlld3BvcnQgICAgIHsgd2lkdGg6IGRldmljZS13aWR0aDsgfQpALW8tdmlld3BvcnQgICAgICB7IHdpZHRoOiBkZXZpY2Utd2lkdGg7IH0KQHZpZXdwb3J0ICAgICAgICAgeyB3aWR0aDogZGV2aWNlLXdpZHRoOyB9Cg=="
	css, _ := base64.StdEncoding.DecodeString(css_b64)
	w.Header().Set("Content-Type", "text/css")
	w.Write(css)
}
