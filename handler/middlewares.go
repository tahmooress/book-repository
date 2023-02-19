package handler

import (
	"net/http"
)

func (h *handler) JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (h *handler) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, ok := r.Header["Token"]
		if !ok || len(t) == 0 {
			h.errorResponse(w, newErr(emptyTokon, nil))

			return
		}

		err := h.au.Verify(t[0])
		if err != nil {
			h.errorResponse(w, serviceErrorMapper(err))

			return
		}

		next.ServeHTTP(w, r)
	})
}
