package postgres

import (
	"context"

	"github.com/irshadpalayadan/job-service-pg-based/model"
	"gorm.io/gorm"
)

type JobListingRepo struct {
	DB *gorm.DB
}

func InitDBRepository(db *gorm.DB) model.IDBRepository {
	return &JobListingRepo{DB: db}
}

func (dbRepo *JobListingRepo) InitTxn(ctx context.Context) (*gorm.DB, func()) {
	dbTxn := dbRepo.DB.Begin()
	return dbTxn, func() {
		if r := recover(); r != nil {
			dbTxn.Rollback()
		}
	}
}

func (dbRepo *JobListingRepo) WithTxn(txnHandler *gorm.DB) model.IDBMethods {
	if txnHandler == nil {
		return nil
	}
	return &JobListingRepo{DB: txnHandler}
}

func (dbRepo *JobListingRepo) GetDBInstanceWithoutTxn() *gorm.DB {
	return dbRepo.DB
}
