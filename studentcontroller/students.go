package studentcontroller

import (
	"CBT/studentsmodel"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controller struct{}

var db *sql.DB

func (control Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var theStudent studentsmodel.Student
		json.NewDecoder(r.Body).Decode(&theStudent)

		var matricNumber string

		err := db.QueryRow("insert into theStudents(matric_number,first_name,last_name,department,gender,email) values($1,$2,$3,$4,$5,$6) RETURNING matric_number", theStudent.Matric_Number, theStudent.First_Name, theStudent.Last_Name, theStudent.Department, theStudent.Gender, theStudent.Email).Scan(&matricNumber)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(matricNumber)

	}
}

func (control Controller) Signin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theStudent studentsmodel.Student
		var checkMatricNumber studentsmodel.Student
		json.NewDecoder(r.Body).Decode(&theStudent)

		row := db.QueryRow("select matric_number from theStudents where matric_number = $1", theStudent.Matric_Number)

		row.Scan(&checkMatricNumber.Matric_Number)

		if theStudent.Matric_Number == checkMatricNumber.Matric_Number {
			fmt.Println("student exist")
		} else {
			fmt.Println("student does not exist")
		}
	}
}

//this controller is for admin and student
func (control Controller) MyProfile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var theStudentInfo studentsmodel.Student
		json.NewDecoder(r.Body).Decode(&theStudentInfo)

		var studentsInfo studentsmodel.Student

		row := db.QueryRow("select matric_number, first_name,last_name,gender,department,operating_system,data_structure,gns,total,average,status from theStudents inner join studentStatus on  matric_number = student_id where matric_number = $1", theStudentInfo.Matric_Number)
		err := row.Scan(&studentsInfo.Matric_Number, &studentsInfo.First_Name, &studentsInfo.Last_Name, &studentsInfo.Gender, &studentsInfo.Department, &studentsInfo.TotalResult.Operating_System, &studentsInfo.TotalResult.Data_Structure, &studentsInfo.TotalResult.Gns, &studentsInfo.TotalResult.Total, &studentsInfo.TotalResult.Average, &studentsInfo.TotalResult.Status)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(studentsInfo)

	}
}
