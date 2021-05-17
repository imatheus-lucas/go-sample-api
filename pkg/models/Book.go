package models

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
