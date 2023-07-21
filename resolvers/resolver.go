package resolvers

import (
	"github.com/irshadpalayadan/job-service-pg-based/model"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	WriteDB model.IDBRepository
	Logger  *zap.Logger
}
