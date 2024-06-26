package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deepakworldphp86/graphql-go/graph/config"
	"github.com/deepakworldphp86/graphql-go/graph/model"

	"gorm.io/gorm"
)

// Setup MySQL Connection
var dbObject *gorm.DB = config.SetupMysqlDatabaseConnection()

// ID is the resolver for the id field.
func (r *categoryResolver) ID(ctx context.Context, obj *model.Category) (*string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// CreatedAt is the resolver for the created_at field.
func (r *categoryResolver) CreatedAt(ctx context.Context, obj *model.Category) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *categoryResolver) UpdatedAt(ctx context.Context, obj *model.Category) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updated_at"))
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	categoryObj := &model.Category{
		CategoryName: input.CategoryName,
		Description:  input.Description,
	}
	// AutoMigrate will create the table if it does not exist

	dbObject.AutoMigrate(&model.Category{})
	result := dbObject.Create(&categoryObj)

	if result.Error != nil {
		panic("failed to insert user")
	}

	fmt.Println("User inserted successfully!")

	//Now Inserting in MongoDB
	client, err := config.SetupMongoDBConnection()

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	categoryCollection := client.Database("shopdb").Collection("categories")

	category := model.CategoryJson{
		CategoryName: input.CategoryName,
		Description:  input.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	insertResult, err := categoryCollection.InsertOne(context.Background(), category)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	return categoryObj, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// GetCategory is the resolver for the getCategory field.
func (r *queryResolver) GetCategory(ctx context.Context, id string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: GetCategory - getCategory"))
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
