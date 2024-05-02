package router

import (
	"net/http"
)

func buildVersion(w http.ResponseWriter) {
	replay := struct {
		Service string `json:"service"`
		Version string `json:"version"`
	}{
		Service: "LightSpeed API",
		Version: "2.0",
	}
	ResponseBuild(w, replay)
	return
}
