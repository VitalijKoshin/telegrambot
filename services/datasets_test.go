package services

import (
	"fmt"
	"time"

	"telegrambot/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testSubsUID = primitive.NewObjectID()
	testSubsID  = primitive.NewObjectID()
	currentTime = time.Time.UTC(time.Time{}).Unix()
)

type FakeSubsModel struct{}

func (s *FakeSubsModel) CreateSubs() models.Subs {
	return models.Subs{}
}

func (s *FakeSubsModel) AddSubs(subs *models.Subs) (*models.Subs, error) {
	return subs, nil
}

func (s *FakeSubsModel) UpdateSubs(subs *models.Subs) (*models.Subs, error) {
	return subs, nil
}

func (s *FakeSubsModel) LoadSubs(ID primitive.ObjectID) (*models.Subs, error) {
	return &models.Subs{}, nil
}

func (s *FakeSubsModel) DeleteSubs(subs *models.Subs) error {
	return nil
}

func (s *FakeSubsModel) FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error) {
	return &[]models.Subs{}, nil
}

func (s *FakeSubsModel) FindNextSubsByTime(fromTime int64, toTime int64) (*[]models.Subs, error) {
	return &[]models.Subs{}, nil
}

func (s *FakeSubsModel) FindUserExistSubs(subs *models.Subs) (*models.Subs, error) {
	return subs, nil
}

type FakeSubsErrorModel struct{}

func (s *FakeSubsErrorModel) CreateSubs() models.Subs {
	return models.Subs{}
}

func (s *FakeSubsErrorModel) AddSubs(subs *models.Subs) (*models.Subs, error) {
	return nil, fmt.Errorf("error add subs")
}

func (s *FakeSubsErrorModel) UpdateSubs(subs *models.Subs) (*models.Subs, error) {
	return nil, fmt.Errorf("error update subs")
}

func (s *FakeSubsErrorModel) LoadSubs(ID primitive.ObjectID) (*models.Subs, error) {
	return nil, fmt.Errorf("error load subs")
}

func (s *FakeSubsErrorModel) DeleteSubs(subs *models.Subs) error {
	return fmt.Errorf("error delete subs")
}

func (s *FakeSubsErrorModel) FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error) {
	return nil, fmt.Errorf("error find subs by uid")
}

func (s *FakeSubsErrorModel) FindNextSubsByTime(fromTime int64, toTime int64) (*[]models.Subs, error) {
	return nil, fmt.Errorf("error find next subs by time")
}

func (s *FakeSubsErrorModel) FindUserExistSubs(subs *models.Subs) (*models.Subs, error) {
	return nil, fmt.Errorf("error find user subs exist")
}

type FakeUserModelSubscribed struct{}

func (u FakeUserModelSubscribed) AddUser(user *models.User) (*models.User, error) {
	return user, nil
}

func (u FakeUserModelSubscribed) Load(id primitive.ObjectID) (*models.User, error) {
	return &models.User{ID: id, TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, nil
}

func (u FakeUserModelSubscribed) UpdateUserOne(user *models.User) (*models.User, error) {
	return user, nil
}

func (u FakeUserModelSubscribed) FindByUsername(username string) (*models.User, error) {
	userF := &models.User{ID: primitive.NewObjectID(), TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}
	return userF, nil
}

func (u FakeUserModelSubscribed) FindByTID(tid int64) (*models.User, error) {
	userF := &models.User{ID: primitive.NewObjectID(), TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}
	return userF, nil
}

func (u FakeUserModelSubscribed) DeleteUser(user *models.User) error {
	return nil
}

func (u FakeUserModelSubscribed) CreateUser() models.User {
	return models.User{}
}

type FakeUserModel struct{}

func (u FakeUserModel) CreateUser() models.User {
	return models.User{}
}

func (u FakeUserModel) Load(id primitive.ObjectID) (*models.User, error) {
	return &models.User{ID: id, TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, nil
}

func (u FakeUserModel) AddUser(user *models.User) (*models.User, error) {
	return user, nil
}

func (u FakeUserModel) UpdateUserOne(user *models.User) (*models.User, error) {
	return user, nil
}

func (u FakeUserModel) FindByUsername(username string) (*models.User, error) {
	return &models.User{ID: primitive.NewObjectID(), TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, nil
}

func (u FakeUserModel) FindByTID(tid int64) (*models.User, error) {
	return &models.User{ID: primitive.NewObjectID(), TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, nil
}

func (u FakeUserModel) DeleteUser(user *models.User) error {
	return nil
}

type FakeUserModelError struct{}

func (u FakeUserModelError) CreateUser() models.User {
	return models.User{}
}

func (u FakeUserModelError) Load(id primitive.ObjectID) (*models.User, error) {
	return &models.User{ID: id, TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, nil
}

func (u FakeUserModelError) AddUser(user *models.User) (*models.User, error) {
	return nil, fmt.Errorf("Error add user")
}

func (u FakeUserModelError) UpdateUserOne(user *models.User) (*models.User, error) {
	return nil, fmt.Errorf("Error update user one")
}

func (u FakeUserModelError) FindByUsername(username string) (*models.User, error) {
	return nil, fmt.Errorf("Error find by username user")
}

func (u FakeUserModelError) FindByTID(tid int64) (*models.User, error) {
	return nil, fmt.Errorf("Error find user by tid")
}

func (u FakeUserModelError) DeleteUser(user *models.User) error {
	return fmt.Errorf("Error delete user")
}
