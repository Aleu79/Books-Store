package store

import (
	"database/sql"
	"practica-go/internal/model"
)

// Esto permite desacoplar la lógica de acceso a datos del resto de la aplicación
type BookStore interface {
	GetAll() ([]*model.Book, error)
	SearchByTitleOrAuthor(book string) ([]*model.Book, error)
	GetByID(id int) (*model.Book, error)
	Exists(id int) (bool, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) (*model.Book, error)
	Delete(id int) error
}

type bookSQL struct {
	db *sql.DB
}

// GetAll obtiene todos los libros de la base de datos
func (s *bookSQL) GetAll() ([]*model.Book, error) {
	q := "SELECT id, title, author FROM books"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []*model.Book

	for rows.Next() {
		b := &model.Book{}
		if err := rows.Scan(&b.ID, &b.Titulo, &b.Autor); err != nil {
			return nil, err
		}
		libros = append(libros, b)
	}

	return libros, nil
}

// SearchByTitleOrAuthor busca libros cuyo título o autor contenga la palabra indicada
func (s *bookSQL) SearchByTitleOrAuthor(book string) ([]*model.Book, error) {
	q := "SELECT id, title, author FROM books WHERE title LIKE ? OR author LIKE ?"

	// Usamos % para permitir coincidencias parciales (ej. "harry" → "Harry Potter")
	rows, err := s.db.Query(q, "%"+book+"%", "%"+book+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []*model.Book

	for rows.Next() {
		b := &model.Book{}
		if err := rows.Scan(&b.ID, &b.Titulo, &b.Autor); err != nil {
			return nil, err
		}
		libros = append(libros, b)
	}
	return libros, nil
}

// GetByID busca un libro por su ID
func (s *bookSQL) GetByID(id int) (*model.Book, error) {
	q := "SELECT id, title, author FROM books WHERE id = ?"

	b := &model.Book{}
	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Titulo, &b.Autor)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Exists verifica si un libro con el ID dado existe en la base de datos
// Usamos SELECT 1 por eficiencia (no se cargan todos los campos)
func (s *bookSQL) Exists(id int) (bool, error) {
	q := "SELECT 1 FROM books WHERE id = ?"
	row := s.db.QueryRow(q, id)

	var exists int
	err := row.Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil // No existe
	}
	if err != nil {
		return false, err // Otro error
	}
	return true, nil // Existe
}

// Create inserta un nuevo libro en la base de datos
func (s *bookSQL) Create(libro *model.Book) (*model.Book, error) {
	q := "INSERT INTO books (title, author) VALUES (?, ?)"
	resp, err := s.db.Exec(q, libro.Titulo, libro.Autor)
	if err != nil {
		return nil, err
	}

	// Obtenemos el ID generado automáticamente por la base de datos
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	libro.ID = int(id)

	return libro, nil
}

// Update actualiza los datos de un libro existente
func (s *bookSQL) Update(id int, libro *model.Book) (*model.Book, error) {
	q := "UPDATE books SET title = ?, author = ? WHERE id = ?"

	_, err := s.db.Exec(q, libro.Titulo, libro.Autor, id)
	if err != nil {
		return nil, err
	}

	libro.ID = id
	return libro, nil
}

// Delete elimina un libro de la base de datos por su ID
func (s *bookSQL) Delete(id int) error {
	q := "DELETE FROM books WHERE id = ?"

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
