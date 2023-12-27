package handlers

import (
	"fmt"
	"time"

	"telegrambot/models"
	weather "telegrambot/pkg/weatherapi"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type weatherClientImplFake struct{}

func (wf *weatherClientImplFake) GetWeatherForecastByLocation(geoUser weather.GeoLocation) (string, error) {
	return "GetWeatherForecastByLocation", nil
}

type weatherClientImplFakeError struct{}

func (wf *weatherClientImplFakeError) GetWeatherForecastByLocation(geoUser weather.GeoLocation) (string, error) {
	return "", fmt.Errorf("GetWeatherForecastByLocation error")
}

type UserServiceFake struct{}

var usrTest = models.User{TID: 777, Username: "username", FirstName: "userFirstName", LastName: "userLastName"}

func (usf UserServiceFake) CreateUser(tid int64, username string, userFirstName string, userLastName string) models.User {
	return usrTest
}

func (usf UserServiceFake) AddUser(user *models.User) (*models.User, error) {
	return &usrTest, nil
}

func (usf UserServiceFake) DeleteUser(user *models.User) (bool, error) {
	return true, nil
}

func (usf UserServiceFake) LoadUserByTID(tid int64) (*models.User, error) {
	return &usrTest, nil
}

func (usf UserServiceFake) LoadUserByUsername(username string) (*models.User, error) {
	return &usrTest, nil
}

func (usf UserServiceFake) UserHasSubscribe(user *models.User) (bool, error) {
	return true, nil
}

func (usf UserServiceFake) SubscribeUser(user *models.User) (bool, error) {
	return true, nil
}

func (usf UserServiceFake) UnSubscribeUserByUserName(userName string) (bool, error) {
	return true, nil
}

func (usf UserServiceFake) UnSubscribeUser(user *models.User) (bool, error) {
	return true, nil
}

type UserServiceFakeError struct{}

func (usf UserServiceFakeError) CreateUser(tid int64, username string, userFirstName string, userLastName string) models.User {
	return usrTest
}

func (usf UserServiceFakeError) AddUser(user *models.User) (*models.User, error) {
	return nil, fmt.Errorf("AddUser error")
}

func (usf UserServiceFakeError) DeleteUser(user *models.User) (bool, error) {
	return true, fmt.Errorf("DeleteUser error")
}

func (usf UserServiceFakeError) LoadUserByTID(tid int64) (*models.User, error) {
	return nil, fmt.Errorf("LoadUserByTID error")
}

func (usf UserServiceFakeError) LoadUserByUsername(username string) (*models.User, error) {
	return nil, fmt.Errorf("LoadUserByUsername error")
}

func (usf UserServiceFakeError) UserHasSubscribe(user *models.User) (bool, error) {
	return true, fmt.Errorf("UserHasSubscribe error")
}

func (usf UserServiceFakeError) SubscribeUser(user *models.User) (bool, error) {
	return true, fmt.Errorf("SubscribeUser error")
}

func (usf UserServiceFakeError) UnSubscribeUserByUserName(userName string) (bool, error) {
	return true, fmt.Errorf("UnSubscribeUserByUserName error")
}

func (usf UserServiceFakeError) UnSubscribeUser(user *models.User) (bool, error) {
	return true, fmt.Errorf("UnSubscribeUser error")
}

type SubsServiceFake struct{}

var (
	subsTest = models.Subs{ChatID: 777, CTime: 777, UID: primitive.NewObjectID(), Latitude: 64, Longitude: 64, Frequency: 86400}
	subss    = []models.Subs{
		subsTest,
	}
)

func (ssf SubsServiceFake) CreateSubs(chatID int64, cTime int64, uid primitive.ObjectID, latitude float64, longitude float64, frequency int64) models.Subs {
	return subsTest
}

func (ssf SubsServiceFake) GetUTCTimeStamp64() int64 {
	return time.Time.UTC(time.Time{}).Unix()
}

func (ssf SubsServiceFake) SetLastTime(ltime int64) int64 {
	return 777
}

func (ssf SubsServiceFake) GetLastTime() int64 {
	return 777
}

func (ssf SubsServiceFake) SaveSubs(subs *models.Subs) (*models.Subs, error) {
	return &subsTest, nil
}

func (ssf SubsServiceFake) LoadSubsByID(id primitive.ObjectID) (*models.Subs, error) {
	return &subsTest, nil
}

func (ssf SubsServiceFake) DeleteSubs(subs *models.Subs) (bool, error) {
	return true, nil
}

func (ssf SubsServiceFake) FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error) {
	return &subss, nil
}

func (ssf SubsServiceFake) UserHasSubscribe(uid primitive.ObjectID, latitude float64, longitude float64) (*[]models.Subs, error) {
	return &subss, nil
}

func (ssf SubsServiceFake) FindSubsBetweenTime(fromTime int64, toTime int64) (*[]models.Subs, error) {
	return &subss, nil
}

func (ssf SubsServiceFake) FindSubsExist(subs *models.Subs) (*models.Subs, error) {
	return &subsTest, nil
}

type ServiceTickerFake struct{}

func (stf ServiceTickerFake) StartTicker() {
}
