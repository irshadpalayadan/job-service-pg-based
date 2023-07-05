package postgres

import (
	"context"
	"time"

	"github.com/irshadpalayadan/job-service-pg-based/model"
)

func (repo *JobListingRepo) GetJobById(ctx context.Context, jobId int) (*model.JobListing, error) {
	jobListing := &model.JobListing{
		Id:           "11",
		Title:        "titley",
		Description:  "Descriptiony",
		Company:      "Companyy",
		Url:          "Urly",
		CreatedBy:    111,
		CreatedAtTs:  time.Now(),
		ModifiedBy:   111,
		ModifiedAtTs: time.Now(),
	}

	return jobListing, nil
}
