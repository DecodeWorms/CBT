package questionsmodel

type Questions struct {
	ID               int    `json:"id"`
	Question_Id      int    `json:"question_id"`
	Operating_System string `json:"operating_system"`
	Data_Structure   string `json:"data_structure"`
	Gns              string `json:"gns"`
}
