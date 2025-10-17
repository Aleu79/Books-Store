package service

import (
	"errors"
	"practica-go/internal/model"
	"strings"
	"unicode"
)

// Trim es un helper para limpiar espacios
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ValidateUser valida un usuario
func ValidateUser(user *model.User) error {
	user.Username = Trim(user.Username)
	user.Email = Trim(user.Email)
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email y password son requeridos")
	}
	if len(user.Password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}
	if !isValidText(user.Username) {
		return errors.New("username contiene caracteres inválidos")
	}
	return nil
}

// ValidateBook valida un libro
func ValidateBook(book *model.Book) error {
	book.Titulo = Trim(book.Titulo)
	book.Autor = Trim(book.Autor)

	if book.Titulo == "" {
		return errors.New("necesitamos el título")
	}
	if len(book.Titulo) < 3 {
		return errors.New("el título es demasiado corto")
	}
	if len(book.Titulo) > 100 {
		return errors.New("el título no puede tener más de 100 caracteres")
	}
	if !isValidText(book.Titulo) {
		return errors.New("el título contiene caracteres inválidos")
	}

	if book.Autor == "" {
		return errors.New("necesitamos el autor")
	}
	if len(book.Autor) < 3 {
		return errors.New("el nombre del autor es demasiado corto")
	}
	if len(book.Autor) > 60 {
		return errors.New("el nombre del autor no puede tener más de 60 caracteres")
	}
	if !isValidText(book.Autor) {
		return errors.New("el nombre del autor contiene caracteres inválidos")
	}

	return nil
}

// isValidText valida caracteres permitidos
func isValidText(text string) bool {
	for _, r := range text {
		switch {
		case unicode.IsLetter(r), unicode.IsNumber(r), unicode.IsSpace(r):
			continue
		case strings.ContainsRune(".,:;!?-'\"()¿¡", r):
			continue
		default:
			return false
		}
	}
	return true
}
