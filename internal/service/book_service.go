package service

import (
	"errors"
	"practica-go/internal/model"
	"practica-go/internal/store"
)

// Service representa la capa de negocio de la aplicación.
// Encapsula la lógica de validación y delega las operaciones
// de persistencia en la capa store (base de datos).
type Service struct {
	store store.Store
}

// New crea una nueva instancia de Service, recibiendo
func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

// GetAllBooks obtiene todos los libros disponibles desde el almacenamiento.
func (s *Service) GetAllBooks() ([]*model.Book, error) {
	return s.store.GetAll()
}

// SearchByTitleOrAuthor busca libros cuyo título o autor contengan
func (s *Service) SearchByTitleOrAuthor(book string) ([]*model.Book, error) {
	return s.store.SearchByTitleOrAuthor(book)
}

// GetByIDBooks obtiene un libro específico según su ID.
func (s *Service) GetByIDBooks(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}

// Exists_Books verifica si existe un libro con el ID dado.
func (s *Service) Exists_Books(id int) (bool, error) {
	return s.store.Exists(id)
}

// Create_Book crea un nuevo libro en la base de datos.
func (s *Service) Create_Book(libro model.Book) (*model.Book, error) {
	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}
	return s.store.Create(&libro)
}

// Update_Book actualiza los datos de un libro existente por ID.
func (s *Service) Update_Book(id int, libro model.Book) (*model.Book, error) {
	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}
	return s.store.Update(id, &libro)
}

// Delete_book elimina un libro existente según su ID.
func (s *Service) Delete_book(id int) error {
	return s.store.Delete(id)
}
