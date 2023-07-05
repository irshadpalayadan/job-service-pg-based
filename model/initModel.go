package model

import (
	"context"

	"gorm.io/gorm"
)

type ICacheRepository interface {
}

type IDBMethods interface {
	IJobListingRepository
}

type IDBRepository interface {
	IDBMethods
	InitTxn(ctx context.Context) (*gorm.DB, func())
	WithTxn(txnHandler *gorm.DB) IDBMethods
	GetDBInstanceWithoutTxn() *gorm.DB
}
