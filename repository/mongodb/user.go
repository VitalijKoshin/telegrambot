package repository

import (
	"context"

	"telegrambot/database/mongodb"
	"telegrambot/internal/config"
	"telegrambot/models"
	"telegrambot/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(cfg *config.Config, mdbClient mongodb.MongoClient) repository.UserRepository {
	collection := mdbClient.GetClient().Database(cfg.MongoDbName).Collection(repository.NameUserCollection)
	return &UserRepositoryImpl{
		userCollection: collection,
		ctx:            context.TODO(),
	}
}

func (u *UserRepositoryImpl) CreateUser() models.User {
	return models.User{}
}

func (u *UserRepositoryImpl) Load(id primitive.ObjectID) (*models.User, error) {
	var userFound models.User
	filter := bson.M{"_id": id}
	err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &userFound, nil
}

func (u *UserRepositoryImpl) AddUser(user *models.User) (*models.User, error) {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) UpdateUserOne(user *models.User) (*models.User, error) {
	filter := bson.M{"_id": user.ID}
	uByte, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}
	var update bson.M
	err = bson.Unmarshal(uByte, &update)
	if err != nil {
		return nil, err
	}
	_, err = u.userCollection.UpdateOne(u.ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) DeleteUser(user *models.User) error {
	_, err := u.userCollection.DeleteOne(u.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var userFound models.User
	filter := bson.M{"username": username}
	err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &userFound, nil
}

func (u *UserRepositoryImpl) FindByTID(tid int64) (*models.User, error) {
	var userFound models.User
	filter := bson.M{"tid": tid}
	err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &userFound, nil
}
