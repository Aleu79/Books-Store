package model

type Book struct {
	ID     int    `json:"id"`
	Titulo string `json:"title"`
	Autor  string `json:"author"`
}
