package service

import (
	"errors"
	"practica-go/internal/model"
	"practica-go/internal/security"
	"practica-go/internal/store"
)

type UserService struct {
	store store.Store
}

func NewUser(s store.Store) *UserService {
	return &UserService{
		store: s,
	}
}

// GetAllUser devuelve todos los usuarios almacenados
func (s *UserService) GetAllUser() ([]*model.User, error) {
	return s.store.UserStorage.GetAllUser()
}

// SearchUserByUserOrEmail busca usuarios por username o email
func (s *UserService) SearchUserByUserOrEmail(term string) ([]*model.User, error) {
	term = Trim(term)
	if term == "" {
		return nil, errors.New("el término no puede estar vacío")
	}
	return s.store.UserStorage.SearchByUserOrEmail(term)
}

// GetUsersByEmailOrUser obtiene un usuario por email o username
func (s *UserService) GetUsersByEmailOrUser(term string) (*model.User, error) {
	term = Trim(term)
	if term == "" {
		return nil, errors.New("el usuario o email no puede estar vacío")
	}

	user, err := s.store.UserStorage.GetByEmailOrUser(term)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}
	return user, nil
}

// ExistsUser verifica si un usuario existe por ID
func (s *UserService) ExistsUser(id int) (bool, error) {
	if id <= 0 {
		return false, errors.New("el id tiene que ser positivo")
	}
	return s.store.UserStorage.Exists(id)
}

// Register crea un nuevo usuario, aplicando validaciones y hash de contraseña
func (s *UserService) Register(user *model.User) (*model.User, error) {
	if err := ValidateUser(user); err != nil {
		return nil, err
	}

	// Comprobar si ya existe usuario con email o username
	existing, _ := s.store.UserStorage.GetByEmailOrUser(user.Username)
	if existing != nil {
		return nil, errors.New("ya existe un usuario con ese username o email")
	}

	// Hashear contraseña y asignar rol por defecto
	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	user.Role = "user"

	created, err := s.store.UserStorage.CreateUser(user)
	if err != nil {
		return nil, err
	}
	created.Password = "" // limpiar contraseña antes de devolver
	return created, nil
}

// Aplica validaciones si se modifican campos y hashea la contraseña si cambio
func (s *UserService) UpdateUser(id int, data *model.User) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("el id debe ser positivo")
	}

	exists, err := s.store.UserStorage.Exists(id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("usuario no encontrado")
	}

	// Validar datos solo si se modifican campos relevantes
	if data.Username != "" || data.Email != "" || data.Password != "" {
		if err := ValidateUser(data); err != nil {
			return nil, err
		}
	}

	// Hashear contraseña si se actualiza
	if data.Password != "" {
		hashed, err := security.HashPassword(data.Password)
		if err != nil {
			return nil, err
		}
		data.Password = hashed
	}

	updated, err := s.store.UserStorage.Update(id, data)
	if err != nil {
		return nil, err
	}
	updated.Password = "" // limpiar contraseña antes de devolver
	return updated, nil
}

// Login valida las credenciales del usuario y devuelve el usuario sin contraseña
func (s *UserService) Login(userOrEmail, password string) (*model.User, error) {
	userOrEmail = Trim(userOrEmail)
	password = Trim(password)
	if userOrEmail == "" || password == "" {
		return nil, errors.New("usuario/email y contraseña son requeridos")
	}

	user, err := s.store.UserStorage.GetByEmailOrUser(userOrEmail)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	if !security.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("contraseña incorrecta")
	}

	user.Password = "" // limpiar password antes de devolver
	return user, nil
}

// DeleteUser elimina un usuario existente según su ID.
func (s *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("el id debe ser positivo")
	}

	exists, err := s.store.UserStorage.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("usuario no encontrado")
	}

	return s.store.UserStorage.Delete(id)
}

// Logout es un marcador: en este service no hace nada.
// Se puede implementar limpieza de tokens o sesiones si se desea.
func (s *UserService) Logout() error {
	return nil
}
