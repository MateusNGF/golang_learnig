package entity

type Person struct {
	Age     int    `json:"age"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Company string `json:"company"`
	Email   string `json:"email"`
}
