package totalresultmodel

type TotalResults struct {
	ID               int    `json:"id"`
	Student_id       string `json:"matric_number"`
	Operating_System int    `json:"operating_system"`
	Data_Structure   int    `json:"data_structure"`
	Gns              int    `json:"gns"`
	Total            int    `json:"total"`
	Average          int    `json:"average"`
	Status           string `json:"status"`
}
