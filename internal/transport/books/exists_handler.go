package books

import (
	"net/http"
	"practica-go/internal/transport"
	"strconv"
	"strings"
)

// Manejo de existencia de libro
func (h *BookHandler) HandleBookExists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		transport.WriteError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/books/exists/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		transport.WriteError(w, http.StatusBadRequest, "id inválido")
		return
	}

	exists, err := h.service.BookExists(id)
	if err != nil {
		transport.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	transport.WriteJSON(w, http.StatusOK, map[string]bool{"exists": exists})
}
