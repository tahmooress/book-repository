package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tahmooress/book-repository/entities"
)

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			h.errorResponse(w, newErr(badRequest, err))

			return
		}

		if e := request.validate(); e != nil {
			h.errorResponse(w, e)
		}

		user, err := h.srv.RegisterUser(r.Context(), &entities.User{
			Name:     request.UserName,
			Email:    request.Email,
			Password: request.Password,
		})
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		h.responseWriter(w, RegisterResponse{
			ID:    user.ID,
			User:  user.Name,
			Email: user.Email,
		},
			http.StatusCreated,
		)
	}
}
