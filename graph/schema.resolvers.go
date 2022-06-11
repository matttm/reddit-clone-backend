package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"reddit-clone-backend/graph/generated"
	"reddit-clone-backend/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string) (*model.Post, error) {
	post := model.Post{
		Title: title,
	}
	err := r.DB.Create(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id float64) (int, error) {
	r.DB.Where("id = ?", id).Delete(&model.Post{})
	return 1, nil
}

func (r *mutationResolver) Login(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	person := model.Person{
		Username: credentials.Username,
	}
	err := r.DB.Create(&person).Error
	if err != nil {
		return nil, err
	}
	validationObject := model.PersonValidationObject{
		Person: &person,
		ValidationErrors: model.PersonValidationErrors{
			Errors: [],
		},
	}
	return &validationObject, nil
}

func (r *mutationResolver) Register(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, body string, id float64, title string) (*model.Post, error) {
	post := model.Post{
		ID:    id,
		Title: title,
		Body:  body,
	}
	err := r.DB.Save(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	hello := "Hello Stranger"
	return hello, nil
}

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
