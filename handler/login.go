package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tahmooress/book-repository/entities"
)

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			h.errorResponse(w, newErr(badRequest, err))

			return
		}

		if e := request.validate(); e != nil {
			h.errorResponse(w, e)

			return
		}

		user, err := h.srv.AuthenticateUser(r.Context(), &entities.User{
			Email:    request.Email,
			Password: request.Password,
		})
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		token, err := h.au.GenerateToken(user.Name)
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		response := LoginResponse{
			Token: token,
		}

		h.responseWriter(w, response, http.StatusOK)
	}
}
