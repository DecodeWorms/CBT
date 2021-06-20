package studentsmodel

import (
	"CBT/totalresultmodel"
)

type Student struct {
	ID            int    `json:"id"`
	Matric_Number string `json:"matric_number"`
	First_Name    string `json:"first_name"`
	Last_Name     string `json:"last_name"`
	Department    string `json:"department"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	TotalResult   totalresultmodel.TotalResults
}
