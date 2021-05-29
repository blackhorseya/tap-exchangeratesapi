package model

// APIResponse declare exchange rates api response struct
type APIResponse struct {
	Success    bool               `json:"success"`
	Timestamp  int                `json:"timestamp"`
	Historical bool               `json:"historical"`
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
}
