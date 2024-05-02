package router

import (
	"encoding/json"
	"net/http"
)

//ResponseBuild response build and send
func ResponseBuild(w http.ResponseWriter, resp interface{}) {
	b, err := json.Marshal(resp)
	if err != nil {
		errResponseBuild(w, err)
		return
	}
	SetHeaders(w)
	w.Write(b)
	return
}

//errResponseBuild error response build and send
func errResponseBuild(w http.ResponseWriter, errresp error) {
	repl := struct {
		Error interface{}
	}{
		Error: struct {
			Message string
		}{
			Message: errresp.Error(),
		},
	}
	b, _ := json.Marshal(repl)
	SetHeaders(w)
	w.Write(b)
	return
}

//FaultResponse send fault 500 with error
func FaultResponse(w http.ResponseWriter, err error) {
	rbody := struct {
		ErrorMessage string
		Status       string
	}{
		ErrorMessage: err.Error(),
		Status:       "Failure",
	}
	b, err := json.Marshal(rbody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(500)
	w.Write(b)
	return
}

//SetHeaders set standard headers
func SetHeaders(w http.ResponseWriter) {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Authorization, Content-Type, X-Content-Type-Options")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}

// IndexHandler route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	replay := struct {
		Version string `json:"version"`
	}{
		Version: "2.0",
	}
	ResponseBuild(w, replay)
}
