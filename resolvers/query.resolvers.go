package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/irshadpalayadan/job-service-pg-based/graph/generated"
	"github.com/irshadpalayadan/job-service-pg-based/graph/model"
)

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: Jobs - jobs"))
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {

	result, _ := r.WriteDB.GetJobById(ctx, 123)
	JobListing := &model.JobListing{
		ID:          result.Id,
		Title:       result.Title,
		Description: result.Description,
		Company:     result.Company,
		URL:         result.Url,
	}
	return JobListing, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
