package transport

import (
	"net/http"
	"strconv"
	"strings"
)

// Manejo de existencia de libro
func (h *BookHandler) HandleBookExists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/books/exists/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id inválido")
		return
	}

	exists, err := h.service.BookExists(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]bool{"exists": exists})
}
