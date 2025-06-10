package user

import "encoding/json"

type Address struct {
	City    string      `json:"addressCity"`
	State   string      `json:"addressState"`
	Country string      `json:"addressCountry"`
	Pincode json.Number `json:"addressPincode"`
}

type User struct {
	Name    string      `json:"userName"`
	Age     json.Number `json:"userAge"`
	Contact string      `json:"userContact"`
	Company string      `json:"userCompany"`
	Address Address     `json:"userAddress"`
}
