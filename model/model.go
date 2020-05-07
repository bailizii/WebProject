package model

type Person struct {
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

