package answercontroller

import (
	"CBT/answermodel"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct{}

var db *sql.DB

func (control Controller) QuestionsAnswers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theAnswers answermodel.Answer
		json.NewDecoder(r.Body).Decode(&theAnswers)

		var answerId int

		err := db.QueryRow("insert into answers(answer_id,operating_system,data_structure,gns) values($1,$2,$3,$4) RETURNING answer_id", theAnswers.AnswerId, theAnswers.Operating_System, theAnswers.Data_Structure, theAnswers.Gns).Scan(&answerId)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(answerId)

	}
}
