package main

import (
	"fmt"
	"practica-go/internal/model"
)

func main() {
	b := model.Book{
		ID:     1,
		Titulo: "El Principito",
		Autor:  "Antoine de Saint-Exupéry",
	}

	fmt.Println(b)
}
