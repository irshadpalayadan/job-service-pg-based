package postgres

import (
	"context"

	"github.com/irshadpalayadan/job-service-pg-based/model"
)

func (repo *JobListingRepo) GetJobById(ctx context.Context, jobId int) (*model.JobListing, error) {
	var jobListingItem model.JobListing
	result := repo.DB.Table("joblist").First(&jobListingItem, "id = ?", "1223")

	if(result.Error != nil) {
		return nil, result.Error
	}

	return &jobListingItem, nil
}
