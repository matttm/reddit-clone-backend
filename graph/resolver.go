package graph

import "reddit-clone-backend/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PersonStore map[string]model.Person
}
