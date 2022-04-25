package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateTodo(todo *Todo) (*primitive.ObjectID, error)
	CreateTodoDetail(todoDetail *TodoDetail) (*primitive.ObjectID, error)
	FindTodoById(id string) (*Todo, error)
	FindTodoDetailById(id string) ([]TodoDetail, error)
	FindTodos() ([]*Todo, error)
	DeleteTodo(id string) error
	DeleteTodoDetail(id string) error
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

type repository struct {
	db *mongo.Database
}

// DeleteTodoDetail implements Repository
func (repo *repository) DeleteTodoDetail(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.db.Collection("todo_details").DeleteMany(nil, bson.D{{"todo_id", mongoId}})
	return err
}

// CreateTodo implements Repository
func (repo *repository) CreateTodo(todoParam *Todo) (*primitive.ObjectID, error) {
	var todo Todo = *todoParam
	ctx := context.Background()
	res, err := repo.db.Collection("todos").InsertOne(ctx, todo)
	id := res.InsertedID.(primitive.ObjectID)
	return &id, err
}

// CreateTodoDetail implements Repository
func (repo *repository) CreateTodoDetail(todoDetail *TodoDetail) (*primitive.ObjectID, error) {
	res, err := repo.db.Collection("todo_details").InsertOne(nil, todoDetail)
	id := res.InsertedID.(primitive.ObjectID)
	return &id, err
}

// DeleteTodo implements Repository
func (repo *repository) DeleteTodo(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.db.Collection("todos").DeleteOne(nil, bson.D{{"_id", mongoId}})
	return err
}

// FindTodoById implements Repository
func (repo *repository) FindTodoById(id string) (*Todo, error) {
	ctx := context.Background()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := repo.db.Collection("todos").Find(ctx, bson.D{{"_id", mongoId}})
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	if res.RemainingBatchLength() < 1 {
		return nil, errors.New("no todo found")
	}
	result := make([]*Todo, 0)
	for res.Next(ctx) {
		var todo Todo
		err := res.Decode(&todo)
		if err != nil {
			return nil, err
		}
		result = append(result, &todo)
	}
	return result[0], nil
}

// FindTodoDetailById implements Repository
func (repo *repository) FindTodoDetailById(id string) ([]TodoDetail, error) {
	ctx := context.Background()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := repo.db.Collection("todo_details").Find(ctx, bson.D{{"todo_id", mongoId}})
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	if res.RemainingBatchLength() < 1 {
		return nil, errors.New("no todo detail found")
	}
	result := make([]TodoDetail, 0)
	for res.Next(ctx) {
		var todoDetail TodoDetail
		err := res.Decode(&todoDetail)
		if err != nil {
			return nil, err
		}
		result = append(result, todoDetail)
	}
	return result, nil
}

// FindTodos implements Repository
func (repo *repository) FindTodos() ([]*Todo, error) {
	ctx := context.Background()
	res, err := repo.db.Collection("todos").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	result := make([]*Todo, 0)
	for res.Next(ctx) {
		var todo Todo
		err := res.Decode(&todo)
		if err != nil {
			return nil, err
		}
		result = append(result, &todo)
	}
	return result, nil
}
