package models

type Todo struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed string `json:"completed"`
}
