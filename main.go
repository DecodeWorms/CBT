package main

import (
	"CBT/questionscontroller"
	"CBT/studentcontroller"
	"CBT/totalresult"
	"blog/connectionDriver"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = connectionDriver.ConnectDB()
	thestudentcontroller := studentcontroller.Controller{}
	theQuestionController := questionscontroller.Controller{}
	theTotalResult := totalresult.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/signup", thestudentcontroller.Signup(db)).Methods("POST")
	router.HandleFunc("/signin", thestudentcontroller.Signin(db)).Methods("GET")
	router.HandleFunc("/myprofile", thestudentcontroller.MyProfile(db)).Methods("GET")

	router.HandleFunc("/osquestions", theQuestionController.TakeOperatingSystemExams(db)).Methods("GET")
	router.HandleFunc("/dataquestions", theQuestionController.TakeDataStructure(db)).Methods("GET")
	router.HandleFunc("/gnsquestions", theQuestionController.TakeGns(db)).Methods("GET")

	router.HandleFunc("/osanswer", theQuestionController.MarkOsScript(db)).Methods("POST")
	router.HandleFunc("/datastructureanswer", theQuestionController.MarkDataScript(db)).Methods("POST")
	router.HandleFunc("/gnsanswer", theQuestionController.MarkGnsAnswer(db)).Methods("POST")

	router.HandleFunc("/setquestions", theQuestionController.SetQuestions(db)).Methods("POST")

	router.HandleFunc("/getresults", theTotalResult.GetTotalResult(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
