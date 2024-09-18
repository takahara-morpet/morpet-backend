package models

type User struct {
	Id      int    `json:"Id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Profile string `json:"profile"`
	Gender  string `json:"gender"`
}
