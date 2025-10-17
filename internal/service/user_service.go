package service

import "practica-go/internal/store"

type UserService struct {
	store store.Store
}

func New(s store.Store) *BookService {
	return &BookService{
		store: s,
	}
}
