package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/ijul/be-monggo/domain/request"
	"github.com/ijul/be-monggo/mongo"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.TaskRepository.
func (t *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := t.database.Collection(t.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

// FetchByUserID implements domain.TaskRepository.
func (t *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return tasks, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		tasks = []domain.Task{}
	}

	return tasks, err
}
