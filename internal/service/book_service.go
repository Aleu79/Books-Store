package service

import (
	"errors"
	"practica-go/internal/model"
	"practica-go/internal/store"
	"strings"
	"unicode"
)

// Service representa la capa de negocio de la aplicación.
// Encapsula la lógica de validación y delega las operaciones
// de persistencia en la capa store (base de datos).
type Service struct {
	store store.Store
}

// New crea una nueva instancia del servicio, recibiendo un store como dependencia.
func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

// GetAllBooks obtiene todos los libros disponibles desde el almacenamiento.
func (s *Service) GetAllBooks() ([]*model.Book, error) {
	return s.store.GetAll()
}

// SearchByTitleOrAuthor busca libros cuyo título o autor contengan el término indicado.
func (s *Service) SearchByTitleOrAuthor(term string) ([]*model.Book, error) {
	term = strings.TrimSpace(term)
	if term == "" {
		return nil, errors.New("el término de búsqueda no puede quedar vacío")
	}
	return s.store.SearchByTitleOrAuthor(term)
}

// GetBookByID obtiene un libro específico según su ID.
func (s *Service) GetBookByID(id int) (*model.Book, error) {
	if id <= 0 {
		return nil, errors.New("el id debe ser positivo")
	}

	book, err := s.store.GetByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("no se encontró el libro con ese id")
	}
	return book, nil
}

// BookExists verifica si existe un libro con el ID dado.
func (s *Service) BookExists(id int) (bool, error) {
	if id <= 0 {
		return false, errors.New("el id debe ser positivo")
	}
	return s.store.Exists(id)
}

// CreateBook crea un nuevo libro en la base de datos, validando sus datos antes.
func (s *Service) CreateBook(libro model.Book) (*model.Book, error) {
	if err := validateBook(&libro); err != nil {
		return nil, err
	}
	return s.store.Create(&libro)
}

// UpdateBook actualiza los datos de un libro existente por ID.
func (s *Service) UpdateBook(id int, libro model.Book) (*model.Book, error) {
	if id <= 0 {
		return nil, errors.New("el id debe ser positivo")
	}

	if err := validateBook(&libro); err != nil {
		return nil, err
	}

	existing, err := s.store.SearchByTitleOrAuthor(libro.Titulo)
	if err != nil {
		return nil, err
	}
	if len(existing) > 0 {
		return nil, errors.New("ya existe un libro con ese título")
	}

	return s.store.Update(id, &libro)
}

// DeleteBook elimina un libro existente según su ID.
func (s *Service) DeleteBook(id int) error {
	if id <= 0 {
		return errors.New("el id debe ser positivo")
	}

	exists, err := s.store.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("no se puede eliminar: el libro no existe")
	}

	return s.store.Delete(id)
}

// validateBook realiza validaciones de negocio sobre los datos del libro.
func validateBook(libro *model.Book) error {
	libro.Titulo = strings.TrimSpace(libro.Titulo)
	libro.Autor = strings.TrimSpace(libro.Autor)

	if libro.Titulo == "" {
		return errors.New("necesitamos el título")
	}
	if len(libro.Titulo) < 3 {
		return errors.New("el título es demasiado corto")
	}
	if len(libro.Titulo) > 100 {
		return errors.New("el título no puede tener más de 100 caracteres")
	}
	if !isValidText(libro.Titulo) {
		return errors.New("el título contiene caracteres inválidos")
	}

	if libro.Autor == "" {
		return errors.New("necesitamos el autor")
	}
	if len(libro.Autor) < 3 {
		return errors.New("el nombre del autor es demasiado corto")
	}
	if len(libro.Autor) > 60 {
		return errors.New("el nombre del autor no puede tener más de 60 caracteres")
	}
	if !isValidText(libro.Autor) {
		return errors.New("el nombre del autor contiene caracteres inválidos")
	}

	return nil
}

// Verifica que el texto solo contenga letras, números, espacios y signos comunes.
// Es una función interna del paquete.
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
