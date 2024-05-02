package router

import (
	"net/http"
	"strings"

	"whapi/rqbag"
)

// RootHdlr route
// lsapi root route
func RootHdlr(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		buildVersion(w)
		return
	}
	rb, errRB := rqbag.BuildReqBag(r.URL)
	if errRB != nil {
		//log.Println("Build RequestBag err: ", errRB)
		errResponseBuild(w, errRB)
		return
	}
	//build response
	//the route is determined by rb.Path
	switch {
	case rb.Path == `partlog`:
		PartlogHandler(w, r, rb)
	default:
		buildVersion(w) //version report response
	}
}
