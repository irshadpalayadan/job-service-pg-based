package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

const (
	CAN_NOT_FETCH_BY_ID   = "CAN_NOT_FETCH_BY_ID"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	GRAPHQL_PARSE_FAILED  = "GRAPHQL_PARSE_FAILED"
)

var errorStatusCode = map[string]int{
	"CAN_NOT_FETCH_BY_ID":   404,
	"INTERNAL_SERVER_ERROR": 500,
}

var internalGqlStatusCode = map[string]int{
	"GRAPHQL_PARSE_FAILED": 403,
}

var internalGqlErrorMsg = map[string]string{
	"GRAPHQL_PARSE_FAILED": "invalid data type provided in the payload",
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

	statusCode, keyPresent = internalGqlStatusCode[_errorCode]
	if keyPresent {
		_logger.Warn("found internal gql error status code")
		return &gqlerror.Error{
			Message: internalGqlErrorMsg[_errorCode],
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

func GqlErrorRespMiddleware(_logger *zap.Logger) func(context.Context, error) *gqlerror.Error {
	return func(_ctx context.Context, _error error) *gqlerror.Error {
		gqlErrorPresenter := graphql.DefaultErrorPresenter(_ctx, _error)
		gqlPresenterErrorCode, _ := gqlErrorPresenter.Extensions["code"]
		formattedStatusCode := fmt.Sprintf("%v", gqlPresenterErrorCode)

		_, isGqlErrCodePresent := internalGqlStatusCode[formattedStatusCode]
		if isGqlErrCodePresent {
			return ConstructErrorResponse(formattedStatusCode, "", _logger)
		}
		return gqlErrorPresenter
	}
}
