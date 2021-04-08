package httperror

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func httpRespondWithError(c echo.Context, err error, logMSg string, status int) *echo.HTTPError {
	eErr := echo.HTTPError{Code: status, Message: logMSg, Internal: err}
	return &eErr
}

func Internal(c echo.Context, err error) *echo.HTTPError {
	return httpRespondWithError(c, err, "Internal server error", http.StatusInternalServerError)
}

func BadRequest(c echo.Context, err error) *echo.HTTPError {
	return httpRespondWithError(c, err, "Bad Request", http.StatusBadRequest)
}
func NotFound(c echo.Context, err error) *echo.HTTPError {
	return httpRespondWithError(c, err, "Not found", http.StatusNotFound)
}
