package services

import (
	"fmt"

	"telegrambot/models"
	"telegrambot/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(tid int64, username string, userFirstName string, userLastName string) models.User
	LoadUserByUsername(username string) (*models.User, error)
	LoadUserByTID(tid int64) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (bool, error)
	UserHasSubscribe(user *models.User) (bool, error)
	SubscribeUser(user *models.User) (bool, error)
	UnSubscribeUser(user *models.User) (bool, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u UserServiceImpl) CreateUser(tid int64, username string, userFirstName string, userLastName string) models.User {
	user := u.userRepository.CreateUser()
	user.TID = tid
	user.Username = username
	user.FirstName = userFirstName
	user.LastName = userLastName
	user.Status = false
	return user
}

func (u UserServiceImpl) AddUser(user *models.User) (*models.User, error) {
	userExists, err := u.userRepository.FindByTID(user.TID)
	if err != nil {
		return nil, fmt.Errorf("user service add user error")
	}

	if userExists != nil && !userExists.ID.IsZero() {
		return userExists, nil
	}

	user.ID = primitive.NewObjectID()
	user.Status = false

	_, err = u.userRepository.AddUser(user)
	if err != nil {
		return nil, fmt.Errorf("user add failed")
	}

	return user, nil
}

func (u UserServiceImpl) DeleteUser(user *models.User) (bool, error) {
	userExists, err := u.userRepository.FindByTID(user.TID)
	if err != nil {
		return false, fmt.Errorf("user service delete user error")
	}

	if userExists.ID.IsZero() {
		return true, nil
	}

	err = u.userRepository.DeleteUser(user)
	if err != nil {
		return false, fmt.Errorf("user service delete user failed")
	}

	return true, nil
}

func (u UserServiceImpl) LoadUserByTID(tid int64) (*models.User, error) {
	userExists, err := u.userRepository.FindByTID(tid)
	if err != nil {
		return nil, fmt.Errorf("load user failed")
	}
	if userExists == nil {
		return nil, fmt.Errorf("load user not exists")
	}
	return userExists, nil
}

func (u UserServiceImpl) LoadUserByUsername(username string) (*models.User, error) {
	userExists, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("load user failed")
	}
	if userExists == nil {
		return nil, fmt.Errorf("load user not exists")
	}
	return userExists, nil
}

func (u UserServiceImpl) UserHasSubscribe(user *models.User) (bool, error) {
	userExists, err := u.userRepository.FindByUsername(user.Username)
	if err != nil {
		return false, fmt.Errorf("check user suscribe failed")
	}
	if userExists == nil {
		return false, nil
	}
	return userExists.Status, nil
}

func (u UserServiceImpl) SubscribeUser(user *models.User) (bool, error) {
	userExists, err := u.userRepository.FindByTID(user.TID)
	if err != nil {
		return false, fmt.Errorf("user suscribe error")
	}

	if userExists.Status {
		return true, nil
	}

	user.Status = true

	_, err = u.userRepository.UpdateUserOne(user)
	if err != nil {
		return false, fmt.Errorf("user subscribeing failed")
	}

	return true, nil
}

func (u UserServiceImpl) UnSubscribeUserByUserName(userName string) (bool, error) {
	userExists, err := u.userRepository.FindByUsername(userName)
	if err != nil {
		return false, fmt.Errorf("user unsuscribe by username error")
	}
	return u.UnSubscribeUser(userExists)
}

func (u UserServiceImpl) UnSubscribeUser(user *models.User) (bool, error) {
	userExists, err := u.userRepository.FindByUsername(user.Username)
	if err != nil {
		return false, fmt.Errorf("user unsuscribe error")
	}
	if userExists.ID.IsZero() {
		return true, fmt.Errorf("user is not exist")
	}
	userExists.Status = false
	_, err = u.userRepository.UpdateUserOne(userExists)
	if err != nil {
		return false, fmt.Errorf("user unsubscribing failed")
	}

	return true, nil
}
