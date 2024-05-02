package rqbag

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"whapi/qparser"
)

// RequestBag struct is populated from request
// Model - table name (to lowcase transformed)
type RequestBag struct {
	ID        string //ID of model
	Path      string //path
	Model     string //model name
	SingleTon bool
	Filter    string
	Top       uint64
	Skip      uint64
	Stat      uint64
	OrderBy   string
}

// BuildReqBag parse path
func BuildReqBag(reqURL *url.URL) (*RequestBag, error) {
	rb, err := parseModel(reqURL)
	if err != nil {
		return nil, err
	}
	//fill query part of RequestBag
	rb.fillOptions(reqURL)
	return rb, nil
}

func parseModel(rawpath *url.URL) (*RequestBag, error) {
	//test valid path - len is 2 max -> model/id
	//pathRaw := strings.TrimLeft(rawpath.Path, config.LSAPIRoutePrefix)
	pathRaw := strings.TrimLeft(rawpath.Path, "/")
	pathRaw = strings.Trim(pathRaw, "/")
	pathslc := strings.Split(pathRaw, "/") // slice of path
	lenpath := len(pathslc)
	if lenpath < 1 { //empty path is wrong
		return nil, fmt.Errorf("wrong path in request (too short) - %s", rawpath.Path)
	}
	if lenpath > 2 { // path is too long
		return nil, fmt.Errorf("wrong path in request (too long) - %s", rawpath.Path)
	}
	rb := new(RequestBag)

	if lenpath == 2 { //we have ID
		rb.ID = pathslc[1]
	}
	rb.Model = strings.ToLower(pathslc[0])
	if strings.HasSuffix(pathslc[0], "(1)") {
		rb.SingleTon = true // set SingleTon if model has suffix (1) like Part(1)
	}
	//we can clear suffix now - remove singleton literal (1)
	resuf := regexp.MustCompile("\\(.*\\)$")
	rb.Model = resuf.ReplaceAllLiteralString(rb.Model, "")
	//create and clear Path
	rb.Path = resuf.ReplaceAllLiteralString(pathslc[0], "")
	return rb, nil
}

func (rb *RequestBag) fillOptions(rp *url.URL) {
	q := rp.Query()
	//fill filter
	filter := q.Get("$filter")
	if len(filter) > 0 {
		rb.Filter = qparser.FilterParse(filter)
	}
	//fill top
	top := q.Get("$top")
	if len(top) > 0 {
		if topUint, err := strconv.ParseUint(top, 10, 64); err == nil {
			rb.Top = topUint
		}
	}
	//fill skip
	skip := q.Get("$skip")
	if len(skip) > 0 {
		if skipUint, err := strconv.ParseUint(skip, 10, 64); err == nil {
			rb.Skip = skipUint
		}
	}
	//fill orderby
	orderby := q.Get("$orderby")
	if len(orderby) > 0 {
		rb.OrderBy = orderby
	}
	//fill stat
	stat := q.Get("$stat")
	if len(stat) > 0 {
		if statUint, err := strconv.ParseUint(stat, 10, 64); err == nil {
			rb.Stat = statUint
		}
	}
}
