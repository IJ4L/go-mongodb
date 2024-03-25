package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	domain "github.com/ijul/be-monggo/domain/request"
	"github.com/ijul/be-monggo/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.UserRepository.
func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

// Fetch implements domain.UserRepository.
func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	// panic("unimplemented")
	return []domain.User{}, nil
}

// GetByEmail implements domain.UserRepository.
func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)

	return user, err
}

// GetByID implements domain.UserRepository.
func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	// panic("unimplemented")
	return domain.User{}, nil
}
