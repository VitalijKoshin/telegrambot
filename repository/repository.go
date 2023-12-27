package repository

import (
	"telegrambot/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const NameUserCollection = "users"

type UserRepository interface {
	CreateUser() models.User
	Load(id primitive.ObjectID) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	UpdateUserOne(user *models.User) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByTID(tid int64) (*models.User, error)
	DeleteUser(user *models.User) error
}

const NameSubsCollection = "subscriptions"

type SubsRepository interface {
	CreateSubs() models.Subs
	AddSubs(subs *models.Subs) (*models.Subs, error)
	UpdateSubs(subs *models.Subs) (*models.Subs, error)
	LoadSubs(ID primitive.ObjectID) (*models.Subs, error)
	DeleteSubs(subs *models.Subs) error
	FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error)
	FindNextSubsByTime(fromTime int64, toTime int64) (*[]models.Subs, error)
	FindUserExistSubs(subs *models.Subs) (*models.Subs, error)
}
