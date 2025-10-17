package users

import (
	"net/http"
	"practica-go/internal/transport"
	"strconv"
	"strings"
)

func (h *UserHandler) HandleBookExists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		transport.WriteError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/users/exists/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		transport.WriteError(w, http.StatusBadRequest, "id invalido")
		return
	}
	exists, err := h.service.ExistsUser(id)
	if err != nil {
		transport.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	transport.WriteJSON(w, http.StatusOK, map[string]bool{"exists": exists})
}
