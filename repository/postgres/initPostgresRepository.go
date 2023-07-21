package postgres

import (
	"context"

	"github.com/irshadpalayadan/job-service-pg-based/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type JobListingRepo struct {
	DB     *gorm.DB
	logger *zap.Logger
}

func InitDBRepository(_db *gorm.DB, _logger *zap.Logger) model.IDBRepository {
	return &JobListingRepo{DB: _db, logger: _logger}
}

func (dbRepo *JobListingRepo) InitTxn(_ctx context.Context) (*gorm.DB, func()) {
	dbTxn := dbRepo.DB.Begin()
	return dbTxn, func() {
		if r := recover(); r != nil {
			dbTxn.Rollback()
		}
	}
}

func (dbRepo *JobListingRepo) WithTxn(_txnHandler *gorm.DB) model.IDBMethods {
	if _txnHandler == nil {
		return nil
	}
	return &JobListingRepo{DB: _txnHandler}
}

func (dbRepo *JobListingRepo) GetDBInstanceWithoutTxn() *gorm.DB {
	return dbRepo.DB
}
