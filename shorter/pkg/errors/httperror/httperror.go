package httperror

import (
	"net/http"

	"github.com/bartam1/mobilfox/shorter/pkg/errors/exterror"
	echo "github.com/labstack/echo/v4"
)

func httpRespondWithError(slug string, err error, status int) *echo.HTTPError {
	eErr := echo.HTTPError{Code: status, Message: slug, Internal: err}
	return &eErr
}

func Internal(slug string, err error) *echo.HTTPError {
	return httpRespondWithError(slug, err, http.StatusInternalServerError)
}

func BadRequest(slug string, err error) *echo.HTTPError {
	return httpRespondWithError(slug, err, http.StatusBadRequest)
}
func NotFound(slug string, err error) *echo.HTTPError {
	return httpRespondWithError(slug, err, http.StatusNotFound)
}
func RespondWithSlugError(err error) *echo.HTTPError {
	slugError, ok := err.(exterror.ExtError)
	if !ok {
		return Internal("Internal server error", err)
	}

	switch slugError.ErrorType() {
	case exterror.ErrorTypeIncorrectInput:
		return BadRequest(slugError.Slug(), slugError)
	case exterror.ErrorTypeRepository:
		return NotFound(slugError.Slug(), slugError)
	default:
		return Internal(slugError.Slug(), slugError)
	}
}
