package store

import (
	"database/sql"
	"practica-go/internal/model"
)

type UserStore interface {
	GetAllUser() ([]*model.User, error)
	SearchByUserOrEmail(user string) ([]*model.User, error)
	CreateUser(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	Exists(id int) (bool, error)
	Update(id int, book *model.User) (*model.User, error)
	Delete(id int) error
}

type userSQL struct {
	db *sql.DB
}

func (s *userSQL) GetAllUser() ([]*model.User, error) {
	return nil, nil
}

func (s *userSQL) SearchByUserOrEmail(user string) ([]*model.User, error) {
	return nil, nil
}

func (s *userSQL) CreateUser(user *model.User) error {
	return nil
}

func (s *userSQL) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}

func (s *userSQL) Exists(id int) (bool, error) {
	return false, nil
}

func (s *userSQL) Update(id int, user *model.User) (*model.User, error) {
	return nil, nil
}

func (s *userSQL) Delete(id int) error {
	return nil
}
