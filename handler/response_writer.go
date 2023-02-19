package handler

import (
	"encoding/json"
	"net/http"
)

func (h *handler) responseWriter(w http.ResponseWriter, body interface{}, status int) {
	b, err := json.Marshal(body)
	if err != nil {
		h.logger.Errorf("marshaling response >> error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(status)

	_, err = w.Write(b)
	if err != nil {
		h.logger.Errorf("writing to response >> error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (h *handler) errorResponse(w http.ResponseWriter, e *Err) {
	if e.reason != nil {
		h.logger.Errorln(e.reason)
	}

	b, err := json.Marshal(e)
	if err != nil {
		h.logger.Errorf("marshaling response >> error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(e.status())

	_, err = w.Write(b)
	if err != nil {
		h.logger.Errorf("writing to response >> error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
