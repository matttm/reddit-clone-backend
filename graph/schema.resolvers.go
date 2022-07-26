package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"reddit-clone-backend/graph/generated"
	"reddit-clone-backend/graph/model"
	"reddit-clone-backend/internal/persons"
	"reddit-clone-backend/internal/posts"
	"reddit-clone-backend/pkg/jwt"
	"strconv"
)

/**
mutation create{
  createPost(post: {title: "test", body: "test body" }){
    post {
      id
      title
      body
    }
  }
}
**/
func (r *mutationResolver) CreatePost(ctx context.Context, post model.PostInput) (*model.PostValidationObject, error) {
	var _post posts.Post
	_post.Title = post.Title
	_post.Body = post.Body
	postID := _post.Save()
	ret := &model.Post{ID: strconv.FormatInt(postID, 10), Title: post.Title, Body: post.Body, Views: 0}
	validationObject := &model.PostValidationObject{
		Post: ret,
		Errors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	return validationObject, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id float64) (int, error) {
	var _post posts.Post
	_post.Id = strconv.FormatFloat(id, 'E', -1, 32)
	_post.Delete()
	return 1, nil
}

func (r *mutationResolver) Login(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	var person persons.Person
	validationObject := model.PersonValidationObject{
		Person: nil,
		Token:  nil,
		ValidationErrors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	person.Username = credentials.Username

	// TODO: hash the pw

	/**
	should i redo the schrma do i can return token?
	**/
	person.Password = credentials.Password

	// TODO: validation checks

	personId := person.Create()
	token, err := jwt.GenerateToken(person.Username)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			err.Error(),
		)
		return &validationObject, err
	}
	ret := &model.Person{ID: strconv.FormatInt(personId, 10), Username: credentials.Username}

	// TODO: send JWT

	validationObject.Person = ret
	validationObject.Token = &token
	return &validationObject, nil
}

func (r *mutationResolver) Register(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	var person persons.Person
	validationObject := model.PersonValidationObject{
		Person: nil,
		Token:  nil,
		ValidationErrors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	person.Username = credentials.Username

	// TODO: hash the pw

	/**
	should i redo the schrma do i can return token?
	**/
	person.Password = credentials.Password

	// TODO: validation checks

	personId := person.Create()
	token, err := jwt.GenerateToken(person.Username)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			err.Error(),
		)
		return &validationObject, err
	}
	ret := &model.Person{ID: strconv.FormatInt(personId, 10), Username: credentials.Username}

	// TODO: send JWT

	validationObject.Person = ret
	validationObject.Token = &token
	return &validationObject, nil

}

func (r *mutationResolver) UpdatePost(ctx context.Context, body string, id float64, title string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello", nil
}

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	dbPersons := persons.GetAll()
	var persons []*model.Person
	for _, v := range dbPersons {
		tmp := &model.Person{
			ID:       v.Id,
			Username: v.Username,
		}
		persons = append(persons, tmp)
	}
	return persons, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

/**
query {
  posts {
    title
    body
    id
  }
}
**/
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts := posts.GetAll()
	var ret []*model.Post
	for _, v := range posts {
		tmp := &model.Post{
			ID:    v.Id,
			Title: v.Title,
			Body:  v.Body,
			Views: v.Views,
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
