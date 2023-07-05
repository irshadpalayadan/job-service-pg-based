package infra

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(needReadWriteDB bool) (writeDB *gorm.DB, readDB *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
		"localhost", "5432", "im_user", "postgres", "im_db", 5, "")

	// Todo: add read write db config

	writeDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	postgresWDB, err := writeDB.DB()
	postgresWDB.SetMaxOpenConns(2)
	postgresWDB.SetMaxIdleConns(1)

	if needReadWriteDB == true {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d search_path=%s",
			"localhost", "5432", "im_user", "postgres", "im_db", 5, "")

		// Todo: add read write db config

		readDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		postgresRDB, _ := readDB.DB()
		postgresRDB.SetMaxOpenConns(2)
		postgresRDB.SetMaxIdleConns(1)
	}

	return writeDB, readDB, err
}
