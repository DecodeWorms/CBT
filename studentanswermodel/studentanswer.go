package studentanswermodel

type StudentAnswer struct {
	ID               int    `json:"id"`
	Answer_id        int    `json:"answer_id"`
	Matric_Number    string `json:"matric_number"`
	Operating_System string `json:"operating_system"`
	Data_Structure   string `json:"data_structure"`
	Gns              string `json:"gns"`
}
