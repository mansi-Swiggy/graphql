package graph

import (
	"example/graph/controllers/user"
	"example/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos        []*model.Todo
	userResolver user.UserRunner
}
