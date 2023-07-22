package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitBootstrap(_db *gorm.DB, _logger *zap.Logger) error {
	//seed.RunSeedInsertion(db, logger)
	return nil
}
