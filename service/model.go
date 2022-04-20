package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

type TodoDetail struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TodoID primitive.ObjectID `json:"todo_id" bson:"todo_id"`
	Item   string             `json:"item" bson:"item"`
}

type TodoAll struct {
	Todo       Todo         `json:"todo"`
	TodoDetail []TodoDetail `json:"todo_detail"`
}
