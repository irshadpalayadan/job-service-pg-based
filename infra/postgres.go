package infra

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB(_needReadWriteDB bool, _logger *zap.Logger) (writeDB *gorm.DB, readDB *gorm.DB, err error) {

	_logger.Info("starting initializing the database")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
		"localhost", "5432", "iam_user", "postgres", "iam_db", 5, "")

	// Todo: add read write db config from env variables

	writeDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_logger.Error("Error while connecting to the write database")
		panic(err)
	}

	postgresWDB, err := writeDB.DB()
	if err != nil {
		_logger.Error("Error while accessing the write db instance ")
		panic(err)
	}

	postgresWDB.SetMaxOpenConns(2)
	postgresWDB.SetMaxIdleConns(1)

	if _needReadWriteDB {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
			"localhost", "5432", "iam_user", "postgres", "iam_db", 5, "")

		// Todo: add read write db config from env variables

		readDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			_logger.Error("Error while connecting to the read database")
			panic(err)
		}

		postgresRDB, err := readDB.DB()
		if err != nil {
			_logger.Error("Error while accessing the read db instance ")
			panic(err)
		}

		postgresRDB.SetMaxOpenConns(2)
		postgresRDB.SetMaxIdleConns(1)
	}

	_logger.Info("initializing the database completed successfully")
	return writeDB, readDB, err
}
