package model

import (
	"fmt"
	"strconv"
	"strings"

	"whapi/rqbag"
)

const (
	PARTLOG = "`ID`,`Cmf`,`PartNumber`,`SupplierCode`," +
		"`LocationID`,`Avail`,`OnOrder`,`Cost`,`MSRP`," +
		"`Bin1`,`ScopeAvail`,`ScopeOnOrder`,`ScopeCost`," +
		"`ScopeMSRP`,`ScopeBin1`,`UpdatedAt`"
)

func buildQuery(rbag *rqbag.RequestBag) (q string, err error) {
	var SelSQL string
	Prefix := "SELECT "
	Suffix := " FROM `" + rbag.Model + "`"
	switch {
	case rbag.Model == "partlog":
		SelSQL = Prefix + PARTLOG + Suffix
	}
	//fmt.Println("SelSQL: ", SelSQL)
	if len(SelSQL) == 0 {
		return q, fmt.Errorf("can't create query %s", rbag.Model)
	}
	//at start, we built 5 elem slice
	// 0 - select part
	// 1 - where part
	// 2 - orderby part
	// 3 - limit part
	// 4 - offset part
	qslc := make([]string, 5)
	qslc[0] = SelSQL // select part
	if len(rbag.Filter) > 0 {
		qslc[1] = "WHERE " + rbag.Filter // where part
	}
	if len(rbag.OrderBy) > 0 {
		qslc[2] = "ORDER BY " + rbag.OrderBy // orderby part
	}
	if rbag.Top > 0 {
		qslc[3] = "LIMIT " + strconv.FormatUint(rbag.Top, 10) // limit part
	}
	if rbag.SingleTon {
		qslc[3] = "LIMIT 1" // singleton limit 1
	}
	if rbag.Skip > 0 {
		qslc[4] = "OFFSET " + strconv.FormatUint(rbag.Skip, 10) // offset part
	}
	//fmt.Println("qslc: ", qslc)
	qslcFin := make([]string, 0)
	for _, el := range qslc {
		if len(el) > 0 {
			qslcFin = append(qslcFin, el)
		}
	}
	q = strings.Join(qslcFin, " ")
	//fmt.Println("q : ", q)
	return
}
