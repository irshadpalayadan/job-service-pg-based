package model

import (
	"context"
	"time"
)

type JobListing struct {
	Id           string    `gorm:"column:id" gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"column:title" json:"title"`
	Description  string    `gorm:"column:description" json:"description"`
	Company      string    `gorm:"column:company" json:"company"`
	Url          string    `gorm:"column:url" json:"url"`
	CreatedBy    int       `gorm:"column:created_by" json:"created_by"`
	CreatedAtTs  time.Time `gorm:"column:created_at_ts" json:"created_at_ts"`
	ModifiedBy   int       `gorm:"column:modified_by" json:"modified_by"`
	ModifiedAtTs time.Time `gorm:"column:modified_at_ts" json:"modified_at_ts"`
}

type IJobListingRepository interface {
	GetJobById(ctx context.Context, jobId string) (*JobListing, error)
}
