package postgres

import (
	"context"

	"github.com/irshadpalayadan/job-service-pg-based/model"
)

func (repo *JobListingRepo) GetJobById(_ctx context.Context, _jobId string) (*model.JobListing, error) {
	repo.logger.Info("started getting job by given id")
	var jobListingItem model.JobListing
	result := repo.DB.Table("public.job_listing").First(&jobListingItem, "id = ?", _jobId)

	if result.Error != nil {
		repo.logger.Error("error while getting job by given id")
		return nil, result.Error
	}

	repo.logger.Info("completed getting job by given id")
	return &jobListingItem, nil
}
