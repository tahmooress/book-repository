package handler

import (
	"fmt"
	"net/http"
)

func (h *handler) SearchBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchKey := r.FormValue("keyword")

		book, err := h.srv.SearchBook(r.Context(), searchKey)
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		fmt.Println(book)
	}
}
