package entity

type Customer struct {
	ID        string  `json:"id"`
	Name      string `json:"name"`
	Balance     int `json:"balance"`
}

type Customers []Customer
