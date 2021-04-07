package exterror

import (
	"github.com/sirupsen/logrus"
)

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnexpected    = ErrorType{"unexpected"}
	ErrorTypeRepository    = ErrorType{"repository"}
	ErrorTypeAuthorization = ErrorType{"authorization"}
)

type ServiceError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s ServiceError) Error() string {
	return s.error
}

func NewUnexpectedError(error string, slug string) ServiceError {
	logrus.WithFields(logrus.Fields{
		"slug": slug,
		"type": ErrorTypeUnexpected.t,
	}).Errorf(error)
	return ServiceError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeUnexpected,
	}
}
func NewRepoError(error string, slug string) ServiceError {
	logrus.WithFields(logrus.Fields{
		"slug": slug,
		"type": ErrorTypeRepository.t,
	}).Errorf(error)
	return ServiceError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeRepository,
	}
}
