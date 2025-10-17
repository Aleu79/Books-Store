package users

import (
	"net/http"
	"practica-go/internal/transport"
	"strings"
)

func (h *UserHandler) HandleSearchUsersOrEmail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		query := strings.TrimSpace(r.URL.Query().Get("q"))
		if query == "" {
			transport.WriteError(w, http.StatusBadRequest, "el término de búsqueda no puede quedar vacío")
			return
		}
		result, err := h.service.SearchUserByUserOrEmail(query)
		if err != nil {
			transport.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		transport.WriteJSON(w, http.StatusOK, map[string]any{"results": result})
	default:
		transport.WriteError(w, http.StatusMethodNotAllowed, "metodo no permitido")
	}

}
