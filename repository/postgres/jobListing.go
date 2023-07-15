package postgres

import (
	"context"
	"fmt"

	"github.com/irshadpalayadan/job-service-pg-based/model"
)

func (repo *JobListingRepo) GetJobById(ctx context.Context, jobId int) (*model.JobListing, error) {
	fmt.Printf("c######### all came GetJobById : ", repo)

	var jobListingItem model.JobListing
	result := repo.DB.Table("public.joblist").First(&jobListingItem, "id = ?", "1223")

	fmt.Printf("result", result)
	fmt.Printf("#######result", result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	return &jobListingItem, nil
}
