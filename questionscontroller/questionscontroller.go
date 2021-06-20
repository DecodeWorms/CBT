package questionscontroller

import (
	"CBT/answermodel"
	"CBT/questionsmodel"
	"CBT/studentanswermodel"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
}

var db *sql.DB

var questions []questionsmodel.Questions

//controller for admin
func (control Controller) SetQuestions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theQuestions questionsmodel.Questions
		json.NewDecoder(r.Body).Decode(&theQuestions)

		var questionId int

		err := db.QueryRow("insert into questions(question_id,operating_system,data_structure,gns) values($1,$2,$3,$4) RETURNING question_id", theQuestions.Question_Id, theQuestions.Operating_System, theQuestions.Data_Structure, theQuestions.Gns).Scan(&questionId)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(questionId)

	}
}

//controller for student
func (control Controller) TakeOperatingSystemExams(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theQuestions questionsmodel.Questions

		questions = []questionsmodel.Questions{}

		rows, err := db.Query("select question_id,operating_system from questions")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err := rows.Scan(&theQuestions.Question_Id, &theQuestions.Operating_System)
			if err != nil {
				log.Fatal(err)
			}

			questions = append(questions, theQuestions)
		}
		json.NewEncoder(w).Encode(questions)
	}
}

//control take action whenever a box is check
func (control Controller) MarkOsScript(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theAnswers studentanswermodel.StudentAnswer
		json.NewDecoder(r.Body).Decode(&theAnswers)

		var comparingResults answermodel.Answer

		var matricNumber string
		var resultMatricNumber string

		var correctScore = 100
		var wrongScore = 0

		err := db.QueryRow("insert into studentAnswers(matric_number,answer_id,operating_system) values($1,$2,$3) RETURNING matric_number", theAnswers.Matric_Number, theAnswers.Answer_id, theAnswers.Operating_System).Scan(&matricNumber)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(theAnswers.Operating_System)

		//code below mark scripts

		row := db.QueryRow("select operating_system from answers where answer_id = $1", theAnswers.Answer_id)

		row.Scan(&comparingResults.Operating_System)

		fmt.Println(comparingResults.Operating_System)

		if theAnswers.Operating_System == comparingResults.Operating_System {
			err := db.QueryRow("insert into results(matric_number,operating_system) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, correctScore).Scan(&resultMatricNumber)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := db.QueryRow("insert into results(matric_number,operating_system) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, wrongScore).Scan(&resultMatricNumber)

			if err != nil {
				log.Fatal(err)
			}
		}

		json.NewEncoder(w).Encode(matricNumber)
		json.NewEncoder(w).Encode(resultMatricNumber)
	}
}

//controller for student

func (control Controller) TakeDataStructure(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theQuestions questionsmodel.Questions

		questions = []questionsmodel.Questions{}

		rows, err := db.Query("select question_id,data_structure from questions")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err := rows.Scan(&theQuestions.Question_Id, &theQuestions.Data_Structure)

			if err != nil {
				log.Fatal(err)
			}

			questions = append(questions, theQuestions)
		}
		json.NewEncoder(w).Encode(questions)
	}

}

//control take action whenever a box is check
func (control Controller) MarkDataScript(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theAnswers studentanswermodel.StudentAnswer
		json.NewDecoder(r.Body).Decode(&theAnswers)

		var matricNumber string
		var comparingResults answermodel.Answer

		var resultMatricNumber string

		var correctScore = 100
		var wrongScore = 0

		err := db.QueryRow("insert into studentAnswers(matric_number,answer_id,data_structure) values($1,$2,$3) RETURNING matric_number", theAnswers.Matric_Number, theAnswers.Answer_id, theAnswers.Data_Structure).Scan(&matricNumber)

		if err != nil {
			log.Fatal(err)
		}

		//comparing and marking script

		row := db.QueryRow("select data_structure from answers where answer_id = $1", theAnswers.Answer_id)

		row.Scan(&comparingResults.Data_Structure)

		fmt.Println(comparingResults.Data_Structure)

		if theAnswers.Data_Structure == comparingResults.Data_Structure {
			err := db.QueryRow("insert into results(matric_number,data_structure) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, correctScore).Scan(&resultMatricNumber)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := db.QueryRow("insert into results(matric_number,data_structure) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, wrongScore).Scan(&resultMatricNumber)

			if err != nil {
				log.Fatal(err)
			}
		}

		json.NewEncoder(w).Encode(matricNumber)
		json.NewEncoder(w).Encode(resultMatricNumber)

	}
}

//controller for student
func (control Controller) TakeGns(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theQuestions questionsmodel.Questions

		questions = []questionsmodel.Questions{}

		rows, err := db.Query("select question_id,gns from questions")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err := rows.Scan(&theQuestions.Question_Id, &theQuestions.Gns)

			if err != nil {
				log.Fatal(err)
			}

			questions = append(questions, theQuestions)
		}
		json.NewEncoder(w).Encode(questions)
	}

}

//control take action whenever a box is check
func (control Controller) MarkGnsAnswer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theAnswers studentanswermodel.StudentAnswer
		json.NewDecoder(r.Body).Decode(&theAnswers)

		var matricNumber string
		var comparingResults answermodel.Answer

		var resultMatricNumber string

		var correctScore = 100
		var wrongScore = 0

		err := db.QueryRow("insert into studentAnswers(matric_number,answer_id,gns) values($1,$2,$3) RETURNING matric_number", theAnswers.Matric_Number, theAnswers.Answer_id, theAnswers.Gns).Scan(&matricNumber)

		if err != nil {
			log.Fatal(err)
		}

		//comparing and marking script

		row := db.QueryRow("select gns from answers where answer_id = $1", theAnswers.Answer_id)

		row.Scan(&comparingResults.Gns)

		fmt.Println(comparingResults.Gns)

		if theAnswers.Gns == comparingResults.Gns {
			err := db.QueryRow("insert into results(matric_number,gns) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, correctScore).Scan(&resultMatricNumber)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := db.QueryRow("insert into results(matric_number,gns) values($1,$2) RETURNING matric_number", theAnswers.Matric_Number, wrongScore).Scan(&resultMatricNumber)

			if err != nil {
				log.Fatal(err)
			}
		}

		json.NewEncoder(w).Encode(matricNumber)
		json.NewEncoder(w).Encode(resultMatricNumber)

	}
}
