package totalresult

import (
	"CBT/resultmodel"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct{}

var db *sql.DB

//admin controller
func (control Controller) GetTotalResult(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var getMatricNumber resultmodel.Results
		json.NewDecoder(r.Body).Decode(&getMatricNumber)

		var operatingSystemTotal int
		var datastructureTotal int
		var gnsTotal int

		operatingSystemTotal = getOsTotal(db, getMatricNumber, w, r)
		datastructureTotal = getDataStructureTotal(db, getMatricNumber, w, r)
		gnsTotal = getGnsTotal(db, getMatricNumber, w, r)

		var total int
		var average int
		var status string

		total = operatingSystemTotal + datastructureTotal + gnsTotal

		average = total / 4

		if average <= 40 {
			status = "fail advice to widthraw"
		} else if average <= 50 {
			status = "pass, you are on probation"
		} else if average <= 60 {
			status = "pass, you are probation"
		} else if average <= 70 {
			status = "Good, you are good to go"
		} else if average <= 80 {
			status = "very good, you are good to go"
		} else if average >= 90 {
			status = "Excellent, what a good performance"
		} else {
			status = "unknow status"
		}

		var thematNumber string

		err := db.QueryRow("insert into studentStatus(student_id,operating_system,data_structure,gns,total,average,status) values($1,$2,$3,$4,$5,$6,$7) RETURNING matric_number", getMatricNumber.Matric_Number, operatingSystemTotal, datastructureTotal, gnsTotal, total, average, status).Scan(&thematNumber)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(thematNumber)

	}

}

func getOsTotal(db *sql.DB, matric_number resultmodel.Results, w http.ResponseWriter, r *http.Request) int {

	var osresultvalue resultmodel.Results

	var osfresult int

	rows := db.QueryRow("select sum(operating_system) from results where matric_number = $1", matric_number.Matric_Number)

	rows.Scan(&osresultvalue.Operating_System)

	osfresult = osresultvalue.Operating_System

	return osfresult

}

func getDataStructureTotal(db *sql.DB, matric_number resultmodel.Results, w http.ResponseWriter, r *http.Request) int {

	var datavalue resultmodel.Results

	var dataresult int

	rows := db.QueryRow("select sum(data_structure) from results where matric_number = $1", matric_number.Matric_Number)

	rows.Scan(&datavalue.Data_Structure)
	dataresult = datavalue.Data_Structure
	return dataresult

}

func getGnsTotal(db *sql.DB, matric_number resultmodel.Results, w http.ResponseWriter, r *http.Request) int {

	var gnsresultvalue resultmodel.Results

	var gnsfresult int

	rows := db.QueryRow("select sum(gns) from results where matric_number = $1", matric_number.Matric_Number)

	rows.Scan(&gnsresultvalue.Gns)
	gnsfresult = gnsresultvalue.Gns
	return gnsfresult

}
