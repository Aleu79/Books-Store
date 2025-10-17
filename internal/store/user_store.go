package store

import (
	"database/sql"
	"practica-go/internal/model"
)

type UserStore interface {
	GetAllUser() ([]*model.User, error)
	SearchByUserOrEmail(user string) ([]*model.User, error)
	GetByEmailOrUser(user string) (*model.User, error)
	Exists(id int) (bool, error)
	CreateUser(user *model.User) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	Delete(id int) error
}

type userSQL struct {
	db *sql.DB
}

func (s *userSQL) GetAllUser() ([]*model.User, error) {
	q := "SELECT id, username, email, role FROM users"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *userSQL) SearchByUserOrEmail(user string) ([]*model.User, error) {
	q := "SELECT id, username, email, role FROM users WHERE username LIKE ? OR email LIKE ?"
	rows, err := s.db.Query(q, "%"+user+"%", "%"+user+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *userSQL) CreateUser(user *model.User) (*model.User, error) {
	q := "INSERT INTO users (username, email, password, role) VALUES(?, ?, ?, ?)"
	resp, err := s.db.Exec(q, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)

	return user, nil
}

func (s *userSQL) GetByEmailOrUser(user string) (*model.User, error) {
	q := "SELECT id, username, email, role FROM users WHERE username = ? OR email = ?"
	row := s.db.QueryRow(q, user, user)

	u := &model.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (s *userSQL) Exists(id int) (bool, error) {
	q := "SELECT 1 FROM users WHERE id = ?"
	row := s.db.QueryRow(q, id)
	var exists int
	err := row.Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *userSQL) Update(id int, user *model.User) (*model.User, error) {
	q := "UPDATE users SET username=?, email=?, role=?, password=? WHERE id=?"
	_, err := s.db.Exec(q, user.Username, user.Email, user.Role, user.Password, id)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (s *userSQL) Delete(id int) error {
	q := "DELETE FROM users WHERE id=?"
	_, err := s.db.Exec(q, id)
	return err
}
