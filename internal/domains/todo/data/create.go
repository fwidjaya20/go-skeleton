package data

// PayloadCreateTodo ...
type PayloadCreateTodo struct {
	Activity    string `json:"activity"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}
