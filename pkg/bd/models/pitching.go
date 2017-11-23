package models

import (
	"errors"
	"log"
	"os"

	"github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

	"github.com/rippinrobr/baseball-databank-db/pkg/db"
)

// Pitching is a model that maps the CSV to a DB Table
type Pitching struct {
	Playerid string  `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
	Yearid   int     `json:"yearID"  csv:"yearID"  db:"yearID"`
	Stint    int     `json:"stint"  csv:"stint"  db:"stint"`
	Teamid   string  `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
	Lgid     string  `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
	W        int     `json:"w"  csv:"W"  db:"W"`
	L        int     `json:"l"  csv:"L"  db:"L"`
	G        int     `json:"g"  csv:"G"  db:"G"`
	Gs       int     `json:"gS"  csv:"GS"  db:"GS"`
	Cg       int     `json:"cG"  csv:"CG"  db:"CG"`
	Sho      int     `json:"sHO"  csv:"SHO"  db:"SHO"`
	Sv       int     `json:"sV"  csv:"SV"  db:"SV"`
	Ipouts   int     `json:"iPouts"  csv:"IPouts"  db:"IPouts"`
	H        int     `json:"h"  csv:"H"  db:"H"`
	Er       int     `json:"eR"  csv:"ER"  db:"ER"`
	Hr       int     `json:"hR"  csv:"HR"  db:"HR"`
	Bb       int     `json:"bB"  csv:"BB"  db:"BB"`
	So       int     `json:"sO"  csv:"SO"  db:"SO"`
	Baopp    float64 `json:"bAOpp"  csv:"BAOpp"  db:"BAOpp"`
	Era      float64 `json:"eRA"  csv:"ERA"  db:"ERA"`
	Ibb      int     `json:"iBB"  csv:"IBB"  db:"IBB"`
	Wp       int     `json:"wP"  csv:"WP"  db:"WP"`
	Hbp      float64 `json:"hBP"  csv:"HBP"  db:"HBP"`
	Bk       int     `json:"bK"  csv:"BK"  db:"BK"`
	Bfp      int     `json:"bFP"  csv:"BFP"  db:"BFP"`
	Gf       int     `json:"gF"  csv:"GF"  db:"GF"`
	R        int     `json:"r"  csv:"R"  db:"R"`
	Sh       int     `json:"sH"  csv:"SH"  db:"SH"`
	Sf       int     `json:"sF"  csv:"SF"  db:"SF"`
	Gidp     int     `json:"gIDP"  csv:"GIDP"  db:"GIDP"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Pitching) GetTableName() string {
	return "pitching"
}

// GetFileName returns the name of the source file the model was created from
func (m *Pitching) GetFileName() string {
	return "Pitching.csv"
}

// GetFilePath returns the path of the source file
func (m *Pitching) GetFilePath() string {
	return "/Users/robertrowe/src/baseballdatabank/core/Pitching.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Pitching) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
	if f == nil {
		return func() error { return nil }, errors.New("nil File")
	}

	return func() error {
		rows := make([]*Pitching, 0)
		numErrors := 0
		err := pfunc(f, &rows)
		if err == nil {
			numrows := len(rows)
			if numrows > 0 {
				log.Println("Pitching ==> Truncating")
				terr := repo.Truncate(m.GetTableName())
				if terr != nil {
					log.Println("truncate err:", terr.Error())
				}

				log.Printf("Pitching ==> Inserting %d Records\n", numrows)
				for _, r := range rows {
					ierr := repo.Insert(m.GetTableName(), r)
					if ierr != nil {
						log.Printf("Insert error: %s\n", ierr.Error())
						numErrors++
					}
				}
			}
			log.Printf("Pitching ==> %d records created\n", numrows-numErrors)
		}

		return err
	}, nil
}
