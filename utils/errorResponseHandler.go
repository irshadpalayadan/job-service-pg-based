package utils

import (
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

const (
	CAN_NOT_FETCH_BY_ID   = "CAN_NOT_FETCH_BY_ID"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
)

var errorStatusCode = map[string]int{
	"CAN_NOT_FETCH_BY_ID":   404,
	"INTERNAL_SERVER_ERROR": 500,
}

func ConstructErrorResponse(_errorCode string, _errorMessage string, _logger *zap.Logger) *gqlerror.Error {
	statusCode, keyPresent := errorStatusCode[_errorCode]
	if keyPresent {
		return &gqlerror.Error{
			Message: _errorMessage,
			Extensions: map[string]interface{}{
				"status_code": statusCode,
				"error_code":  _errorCode,
				"timestamp":   time.Now()},
		}
	}

	_logger.Warn("requested error status code doesn't present")
	return &gqlerror.Error{
		Message: _errorMessage,
		Extensions: map[string]interface{}{
			"status_code": errorStatusCode[INTERNAL_SERVER_ERROR],
			"error_code":  INTERNAL_SERVER_ERROR,
			"timestamp":   time.Now()},
	}
}
