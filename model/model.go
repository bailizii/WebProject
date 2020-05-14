package model

type Person struct {
	ID			string    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	//Lastname    string `json:"lastname,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

