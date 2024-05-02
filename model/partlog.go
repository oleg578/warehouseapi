package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"whapi/rqbag"
)

type Partlog struct {
	ID           string  `json:"ID"`
	Cmf          string  `json:"Cmf"`
	PartNumber   string  `json:"PartNumber"`
	SupplierCode string  `json:"SupplierCode"`
	LocationID   int64   `json:"LocationID"`
	Avail        int64   `json:"Avail"`
	OnOrder      int64   `json:"OnOrder"`
	Cost         float64 `json:"Cost"`
	MSRP         float64 `json:"MSRP"`
	Bin1         string  `json:"Bin1"`
	ScopeAvail   bool    `json:"ScopeAvail"`
	ScopeOnOrder bool    `json:"ScopeOnOrder"`
	ScopeCost    bool    `json:"ScopeCost"`
	ScopeMSRP    bool    `json:"ScopeMSRP"`
	ScopeBin1    bool    `json:"ScopeBin1"`
	UpdatedAt    string  `json:"UpdatedAt"`
}

type Stat struct {
	UpdateDate string `json:"UpdateDate"`
	CountItems int64  `json:"CountItems"`
}

// GetByID get  by ID
func (p *Partlog) GetByID(db *sql.DB, rbag *rqbag.RequestBag) error {
	var (
		PSQL = "SELECT " +
			PARTLOG + " " +
			"FROM `partlog` WHERE `ID`=?"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := db.QueryRowContext(ctx, PSQL, rbag.ID).Scan(
		&p.ID,
		&p.Cmf,
		&p.PartNumber,
		&p.SupplierCode,
		&p.LocationID,
		&p.Avail,
		&p.OnOrder,
		&p.Cost,
		&p.MSRP,
		&p.Bin1,
		&p.ScopeAvail,
		&p.ScopeOnOrder,
		&p.ScopeCost,
		&p.ScopeMSRP,
		&p.ScopeBin1,
		&p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("item fetch error: %s", err.Error())
	}
	return nil
}

func GetFilteredParts(db *sql.DB, rbag *rqbag.RequestBag) (items []Partlog, err error) {
	query, errq := buildQuery(rbag)
	if errq != nil {
		err = errq
		return
	}

	switch {
	case rbag.SingleTon:
		return getPartlogSingle(db, query)
	default:
		return getPartlogSlc(db, query)
	}
}

func getPartlogSingle(db *sql.DB, query string) (items []Partlog, err error) {
	p := Partlog{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = db.QueryRowContext(ctx, query).Scan(
		&p.ID,
		&p.Cmf,
		&p.PartNumber,
		&p.SupplierCode,
		&p.LocationID,
		&p.Avail,
		&p.OnOrder,
		&p.Cost,
		&p.MSRP,
		&p.Bin1,
		&p.ScopeAvail,
		&p.ScopeOnOrder,
		&p.ScopeCost,
		&p.ScopeMSRP,
		&p.ScopeBin1,
		&p.UpdatedAt)
	if err != nil {
		return items, fmt.Errorf("single item fetch error: %s", err.Error())
	}
	items = append(items, p)
	return
}

func getPartlogSlc(db *sql.DB, query string) (items []Partlog, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rows, errR := db.QueryContext(ctx, query)
	if errR != nil {
		err = errR
		return
	}
	for rows.Next() {
		p := Partlog{}
		if err = rows.Scan(
			&p.ID,
			&p.Cmf,
			&p.PartNumber,
			&p.SupplierCode,
			&p.LocationID,
			&p.Avail,
			&p.OnOrder,
			&p.Cost,
			&p.MSRP,
			&p.Bin1,
			&p.ScopeAvail,
			&p.ScopeOnOrder,
			&p.ScopeCost,
			&p.ScopeMSRP,
			&p.ScopeBin1,
			&p.UpdatedAt); err != nil {
			return
		}
		items = append(items, p)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func GetStat(db *sql.DB, rbag *rqbag.RequestBag) (sts []Stat, err error) {
	var (
		query = "SELECT DISTINCT COUNT(PartNumber) AS `CountItems`," +
			"DATE(`UpdatedAt`) AS `UpdateDate` " +
			"FROM `warehouse`.`partlog` " +
			"GROUP BY DATE(`UpdatedAt`) " +
			"ORDER BY DATE(`UpdatedAt`) DESC LIMIT ?"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rows, errR := db.QueryContext(ctx, query, rbag.Stat)
	if errR != nil {
		err = errR
		return
	}
	for rows.Next() {
		s := Stat{}
		if err = rows.Scan(&s.CountItems, &s.UpdateDate); err != nil {
			return
		}
		sts = append(sts, s)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
