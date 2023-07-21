package seed

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RunSeedInsertion(_db *gorm.DB, _logger *zap.Logger) error {

	// result := db.Table('public.job_listing').Create()
	// if result.Error != nil {
	// 	logger.Error('error while inserting the seed data')
	// 	return nil
	// }

	// logger.Info('Seed data inserted successfully')
	return nil
}
