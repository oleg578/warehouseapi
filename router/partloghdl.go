package router

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"whapi/config"
	"whapi/model"
	"whapi/rqbag"
)

func PartlogHandler(w http.ResponseWriter, r *http.Request, rb *rqbag.RequestBag) {
	switch {
	case len(rb.ID) > 0:
		partlogSingleGet(w, rb)
	case rb.Stat > 0:
		partlogStat(w, rb)
	default:
		partlogFilterGet(w, rb)
	}
}

func partlogStat(w http.ResponseWriter, rb *rqbag.RequestBag) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		errResponseBuild(w, err)
		return
	}
	defer db.Close()
	stats, errStat := model.GetStat(db, rb)
	if errStat != nil {
		errResponseBuild(w, errStat)
		return
	}
	ResponseBuild(w, stats)
}

func partlogSingleGet(w http.ResponseWriter, rb *rqbag.RequestBag) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		errResponseBuild(w, err)
		return
	}
	defer db.Close()

	part := model.Partlog{}
	if errGp := part.GetByID(db, rb); errGp != nil {
		errResponseBuild(w, errGp)
		return
	}
	repl := struct {
		Part model.Partlog
	}{
		Part: part,
	}
	ResponseBuild(w, repl)
}

// orderFilterGet get orders array
func partlogFilterGet(w http.ResponseWriter, rb *rqbag.RequestBag) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		errResponseBuild(w, err)
		return
	}
	defer db.Close()
	parts, errParts := model.GetFilteredParts(db, rb)
	if errParts != nil {
		errResponseBuild(w, errParts)
		return
	}
	ResponseBuild(w, parts)
}
