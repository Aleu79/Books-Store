package transport

import (
	"encoding/json"
	"net/http"
	"practica-go/internal/model"
	"practica-go/internal/service"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{service: s}
}

// Manejo de todos los libros
func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		libros, err := h.service.GetAllBooks()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"books": libros})
	case http.MethodPost:
		var libro model.Book
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			writeError(w, http.StatusBadRequest, "input inválido")
			return
		}
		created, err := h.service.CreateBook(libro)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, map[string]any{"book": created})
	default:
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

// Manejo de libro por ID
func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id inválido")
		return
	}

	switch r.Method {
	case http.MethodGet:
		libro, err := h.service.GetBookByID(id)
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"book": libro})

	case http.MethodPut:
		var libro model.Book
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			writeError(w, http.StatusBadRequest, "input inválido")
			return
		}
		updated, err := h.service.UpdateBook(id, libro)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"book": updated})

	case http.MethodDelete:
		if err := h.service.DeleteBook(id); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusNoContent, map[string]string{"message": "el libro fue eliminado"})
	default:
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}
