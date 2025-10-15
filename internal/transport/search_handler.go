package transport

import (
	"net/http"
	"strings"
)

// Manejo de búsqueda de libros
func (h *BookHandler) HandleSearchBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		writeError(w, http.StatusBadRequest, "el término de búsqueda no puede quedar vacío")
		return
	}
	results, err := h.service.SearchByTitleOrAuthor(query)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"results": results})
}
