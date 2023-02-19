package handler

import (
	"net/http"
)

func (h *handler) SearchBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchKey := r.FormValue("keyword")

		books, err := h.srv.SearchBook(r.Context(), searchKey)
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		h.responseWriter(w, SearchBookResult{Books: books}, http.StatusOK)
	}
}
