package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"reddit-clone-backend/graph/generated"
	"reddit-clone-backend/graph/model"
	auth "reddit-clone-backend/internal/auth"
	"reddit-clone-backend/internal/persons"
	"reddit-clone-backend/internal/posts"
	"reddit-clone-backend/pkg/jwt"
	"strconv"
)

/**
  * @function CreatePost
  * @description adds a post to the db the connected db
  *
  * @return a post validation object
*/
func (r *mutationResolver) CreatePost(ctx context.Context, post model.PostInput) (*model.PostValidationObject, error) {
	// determine user suthenticity
	person := auth.ForContext(ctx)
	if person == nil {
		return &model.PostValidationObject{}, fmt.Errorf("access denied")
	}
	// create post
	var _post posts.Post
	_post.Person = person
	_post.Title = post.Title
	_post.Body = post.Body
	postID, err := _post.Save()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	ret := &model.Post{ID: strconv.FormatInt(postID, 10), Title: post.Title, Body: post.Body, Views: 0}
	validationObject := &model.PostValidationObject{
		Post: ret,
		Errors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	return validationObject, nil
}

/**
  * @function DeletePost
  * @description remove a post to the db the connected db
  *
  * @return the deleted post id
*/
func (r *mutationResolver) DeletePost(ctx context.Context, id float64) (int, error) {
	var _post posts.Post
	_post.Id = strconv.FormatFloat(id, 'E', -1, 32)
	_id, err :=_post.Delete()
	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}
	return int(_id), nil
}

/**
  * @function Login
  * @description login with given credentials
  *
  * @return a person validation object
*/
func (r *mutationResolver) Login(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	log.Printf("Attempting login for %s", credentials.Username)
	validationObject := model.PersonValidationObject{
		Person: nil,
		Token:  nil,
		ValidationErrors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	// check whether a username was orovided
	if credentials.Username == "" {
		log.Panicf("Error: username was not provided")
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			"Username was not provided",
		)
		return &validationObject, nil
	}
	// determine authenticity of credentials
	isAuth, err := persons.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	log.Printf("User %s authenticated", credentials.Username)
	//
	// if user cannot be authenticated, throw an error
	if !isAuth {
		log.Panicf("Error: credentials do not match user")
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			"Username or password is incorrect",
		)
		return &validationObject, nil
	}
	// generate token
	token, err := jwt.GenerateToken(credentials.Username)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			err.Error(),
		)
		return &validationObject, err
	}
	//
	// sinxe user was authenticated without error, get the user id to return
	personId, err := persons.GetUserIdByUsername(credentials.Username)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			err.Error(),
		)
		return &validationObject, err
	}
	ret := &model.Person{ID: fmt.Sprint(personId), Username: credentials.Username}

	validationObject.Person = ret
	validationObject.Token = &token
	return &validationObject, nil
}

/**
  * @function Register
  * @description register a user with given credentials
  *
  * @return a person validation object
*/
func (r *mutationResolver) Register(ctx context.Context, credentials model.Credentials) (*model.PersonValidationObject, error) {
	log.Printf("Attempting registration for %s", credentials.Username)
	var person persons.Person
	validationObject := model.PersonValidationObject{
		Person: nil,
		Token:  nil,
		ValidationErrors: &model.ValidationErrors{
			Errors: nil,
		},
	}
	person.Username = credentials.Username
	person.Password = credentials.Password

	// TODO: validation checks
	// check username length
	// check password complexity

	personId, err := person.Create()
	if err != nil {
		log.Panicf("Error: %s", err.Error())
		validationObject.ValidationErrors.Errors = append(
			validationObject.ValidationErrors.Errors,
			err.Error(),
		)
		return &validationObject, err
	}
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

	validationObject.Person = ret
	validationObject.Token = &token
	return &validationObject, nil

}

/**
  * @function UpdatePost
  * @description updates a post
  *
  * @return the updated post
*/
func (r *mutationResolver) UpdatePost(ctx context.Context, body string, id float64, title string) (*model.Post, error) {
	var _post posts.Post
	_post.Id = fmt.Sprintf("%f", id)
	_post.Title = title
	_post.Body = body
	postID, err := _post.Update()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	ret := &model.Post{
		ID:        strconv.FormatInt(postID, 10),
		Title:     _post.Title,
		Body:      _post.Body,
		Views:     _post.Views,
		CreatedAt: _post.CreatedAt,
		UpdatedAt: _post.UpdatedAt,
	}

	return ret, nil
}

/**
  * @function RefreshToken
  * @description refreshes the application's auth token
  *
  * @return the new token
*/
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

/**
  * @function Persons
  * @description gets all persons from db
  *
  * @return an array of persons
*/
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	dbPersons, err := persons.GetAll()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	var persons []*model.Person
	for _, v := range dbPersons {
		tmp := &model.Person{
			ID:        v.Id,
			Username:  v.Username,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		persons = append(persons, tmp)
	}
	return persons, nil
}

/**
  * @function Post
  * @description a post from db
  *
  * @return the post
*/
func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	dbPost := posts.Get(id)
	post := &model.Post{
		ID:    dbPost.Id,
		Title: dbPost.Title,
		Body:  dbPost.Body,
		Person: &model.Person{
			ID:        dbPost.Person.Id,
			Username:  dbPost.Person.Username,
			CreatedAt: dbPost.Person.CreatedAt,
		},
		Views:     dbPost.Views,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
	}
	return post, nil
}

/**
  * @function Posts
  * @description gets all posts from db
  *
  * @return an array of posts
*/
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts := posts.GetAll()
	var ret []*model.Post
	for _, v := range posts {
		tmp := &model.Post{
			ID:    v.Id,
			Title: v.Title,
			Body:  v.Body,
			Person: &model.Person{
				ID:       v.Person.Id,
				Username: v.Person.Username,
			},
			Views:     v.Views,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

/**
  * @function IsAuthenticated
  * @description determines whether user is authenticated
  *
  * @return a boolean of whether user is authenticated
*/
func (r *queryResolver) IsAuthenticated(ctx context.Context) (*model.PersonValidationObject, error) { // determine user suthenticity
	person := auth.ForContext(ctx)
	if person == nil {
		return &model.PersonValidationObject{
			ValidationErrors: &model.ValidationErrors{
				Errors: []string{"Unauthenticated"},
			},
		}, nil
	}
	return &model.PersonValidationObject{
		Person: &model.Person{
			ID:       person.Id,
			Username: person.Username,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
