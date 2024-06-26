package graph

import "github.com/deepakworldphp86/graphql-go/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos        []*model.Todo
	category     []*model.Category
	new_category []*model.NewCategory
}
