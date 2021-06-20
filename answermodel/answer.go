package answermodel

type Answer struct {
	ID               int    `json:"id"`
	AnswerId         int    `json:"answerId"`
	Operating_System string `json:"operating_system"`
	Data_Structure   string `json:"data_structure"`
	Gns              string `json:"gns"`
}
