package service

import (
	"errors"
	"practica-go/internal/model"
	"practica-go/internal/store"
)

// Service representa la capa de negocio de la aplicación.
// Encapsula la lógica de validación y delega las operaciones
// de persistencia en la capa store (base de datos).
type BookService struct {
	store store.Store
}

// NewBook crea una nueva instancia del servicio, recibiendo un store como dependencia.
func NewBook(s store.Store) *BookService {
	return &BookService{
		store: s,
	}
}

// GetAllBooks obtiene todos los libros disponibles desde el almacenamiento.
func (s *BookService) GetAllBooks() ([]*model.Book, error) {
	return s.store.BookStorage.GetAll()
}

// SearchByTitleOrAuthor busca libros cuyo título o autor contengan el término indicado.
func (s *BookService) SearchBookByTitleOrAuthor(term string) ([]*model.Book, error) {
	term = Trim(term)
	if term == "" {
		return nil, errors.New("el término de búsqueda no puede quedar vacío")
	}
	return s.store.BookStorage.SearchByTitleOrAuthor(term)
}

// GetBookByID obtiene un libro específico según su ID.
func (s *BookService) GetBookByID(id int) (*model.Book, error) {
	if id <= 0 {
		return nil, errors.New("el id debe ser positivo")
	}

	book, err := s.store.BookStorage.GetByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("no se encontró el libro con ese id")
	}
	return book, nil
}

// BookExists verifica si existe un libro con el ID dado.
func (s *BookService) BookExists(id int) (bool, error) {
	if id <= 0 {
		return false, errors.New("el id debe ser positivo")
	}
	return s.store.BookStorage.Exists(id)
}

// CreateBook crea un nuevo libro en la base de datos, validando sus datos antes.
func (s *BookService) CreateBook(libro *model.Book) (*model.Book, error) {
	if err := ValidateBook(libro); err != nil {
		return nil, err
	}
	return s.store.BookStorage.Create(libro)
}

// UpdateBook actualiza los datos de un libro existente por ID.
func (s *BookService) UpdateBook(id int, libro *model.Book) (*model.Book, error) {
	if id <= 0 {
		return nil, errors.New("el id debe ser positivo")
	}

	if err := ValidateBook(libro); err != nil {
		return nil, err
	}

	existing, err := s.store.BookStorage.SearchByTitleOrAuthor(libro.Titulo)
	if err != nil {
		return nil, err
	}
	if len(existing) > 0 {
		return nil, errors.New("ya existe un libro con ese título")
	}

	return s.store.BookStorage.Update(id, libro)
}

// DeleteBook elimina un libro existente según su ID.
func (s *BookService) DeleteBook(id int) error {
	if id <= 0 {
		return errors.New("el id debe ser positivo")
	}

	exists, err := s.store.BookStorage.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("no se puede eliminar: el libro no existe")
	}

	return s.store.BookStorage.Delete(id)
}
