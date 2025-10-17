package books

import (
	"net/http"
	"practica-go/internal/transport"
	"strings"
)

// Manejo de búsqueda de libros
func (h *BookHandler) HandleSearchBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		transport.WriteError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		transport.WriteError(w, http.StatusBadRequest, "el término de búsqueda no puede quedar vacío")
		return
	}
	results, err := h.service.SearchBookByTitleOrAuthor(query)
	if err != nil {
		transport.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	transport.WriteJSON(w, http.StatusOK, map[string]any{"results": results})
}
