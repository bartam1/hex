package httperror

import (
	"net/http"

	"github.com/bartam1/mobilfox/shorter/pkg/logs/httplog"
	echo "github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func httpRespondWithError(c echo.Context, err error, logMSg string, status int) {
	logrus.WithFields(logrus.Fields{
		"Internal": err.Error(),
	}).Errorf(logMSg)
	eErr := echo.HTTPError{Code: status, Message: logMSg, Internal: err}
	httplog.ErrorHandler(&eErr, c)
}

func InternalError(c echo.Context, err error) {
	httpRespondWithError(c, err, "Internal server error", http.StatusInternalServerError)
}

func BadRequest(c echo.Context, err error) {
	httpRespondWithError(c, err, "Bad Request", http.StatusInternalServerError)
}
