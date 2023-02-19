package handler

import (
	"errors"
	"net/http"

	"github.com/tahmooress/book-repository/entities"
)

const (
	badRequest    = "bad request format"
	emptyTokon    = "empty token"
	emptyPassword = "empty password"
	emptyUser     = "empty user"
	emailMalform  = "malform email"
	internal      = "something goes wrong, try again!"
)

const (
	internalCode = iota
	unAuthorizeCode
	duplicateUserCode
	invalidPsswordCode
	notFoundCode
	emptyTokonCode
	emptyPasswordCode
	emptyUserCode
	emailMalformCode
	badRequestCode
)

type Type string

type Err struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	reason     error
	httpStatus int
}

func (e *Err) status() int {
	return e.httpStatus
}

func newErr(message string, reason error) *Err {
	switch message {
	case emailMalform:
		return &Err{
			Message:    message,
			Code:       emailMalformCode,
			reason:     reason,
			httpStatus: http.StatusBadRequest,
		}
	case emptyPassword:
		return &Err{
			Message:    message,
			Code:       emptyPasswordCode,
			reason:     reason,
			httpStatus: http.StatusBadRequest,
		}
	case emptyUser:
		return &Err{
			Message:    message,
			Code:       emptyUserCode,
			reason:     reason,
			httpStatus: http.StatusBadRequest,
		}
	case emptyTokon:
		return &Err{
			Message:    message,
			Code:       emptyTokonCode,
			reason:     reason,
			httpStatus: http.StatusUnauthorized,
		}
	case badRequest:
		return &Err{
			Message:    message,
			Code:       badRequestCode,
			reason:     reason,
			httpStatus: http.StatusBadRequest,
		}
	default:
		return &Err{
			Message:    internal,
			Code:       internalCode,
			reason:     reason,
			httpStatus: http.StatusInternalServerError,
		}
	}
}

func serviceErrorMapper(err error) *Err {
	var e *entities.Err

	if !errors.As(err, &e) {
		return &Err{
			Message:    internal,
			Code:       internalCode,
			reason:     err,
			httpStatus: http.StatusInternalServerError,
		}
	}

	switch e.Type {
	case entities.Internal:
		return &Err{
			Code:       internalCode,
			Message:    e.Message(),
			reason:     e.Reason,
			httpStatus: http.StatusInternalServerError,
		}
	case entities.DuplicateUser:
		return &Err{
			Code:       duplicateUserCode,
			Message:    e.Message(),
			reason:     e.Reason,
			httpStatus: http.StatusBadRequest,
		}
	case entities.InvalidPssword:
		return &Err{
			Code:       invalidPsswordCode,
			Message:    e.Message(),
			reason:     e.Reason,
			httpStatus: http.StatusUnauthorized,
		}
	case entities.NotFound:
		return &Err{
			Code:       notFoundCode,
			Message:    e.Message(),
			reason:     e.Reason,
			httpStatus: http.StatusNotFound,
		}
	case entities.UnAuthorize:
		return &Err{
			Code:       unAuthorizeCode,
			Message:    e.Message(),
			reason:     e.Reason,
			httpStatus: http.StatusUnauthorized,
		}
	default:
		return &Err{
			Message:    internal,
			Code:       internalCode,
			httpStatus: http.StatusInternalServerError,
		}
	}
}
