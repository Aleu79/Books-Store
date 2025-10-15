package store

import "database/sql"

// Store centraliza el acceso a los distintos repositorios
type Store struct {
	db          *sql.DB
	BookStorage BookStore
	UserStorage UserStore
}

// New crea una instancia de Store con todas las dependencias inicializadas
func New(db *sql.DB) *Store {
	return &Store{
		db:          db,
		BookStorage: &bookSQL{db: db},
		UserStorage: &userSQL{db: db},
	}
}
