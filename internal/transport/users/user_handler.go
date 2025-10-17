package users

import (
	"encoding/json"
	"net/http"
	"practica-go/internal/model"
	"practica-go/internal/service"
	"practica-go/internal/transport"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func NewHandlerUser(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.service.GetAllUser()
		if err != nil {
			transport.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		transport.WriteJSON(w, http.StatusOK, map[string]any{"users": users})
	case http.MethodPost:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			transport.WriteError(w, http.StatusBadRequest, "input no valido")
			return
		}
		created, err := h.service.Register(&user)
		if err != nil {
			transport.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		transport.WriteJSON(w, http.StatusOK, map[string]any{"user": created})
	default:
		transport.WriteError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

func (h *UserHandler) HandleUserByUserOrEmail(w http.ResponseWriter, r *http.Request) {
	userStr := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case http.MethodGet:
		user, err := h.service.GetUsersByEmailOrUser(userStr)
		if err != nil {
			transport.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		if user == nil {
			transport.WriteError(w, http.StatusNotFound, "usuario no encontrado")
			return
		}

		transport.WriteJSON(w, http.StatusOK, map[string]any{"user": user})

	default:
		transport.WriteError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}
