package services

import (
	"fmt"
	"time"

	"telegrambot/models"
	"telegrambot/repository"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubsService interface {
	CreateSubs(chatID int64, cTime int64, uid primitive.ObjectID, latitude float64, longitude float64, frequency int64) models.Subs
	GetUTCTimeStamp64() int64
	SetLastTime(ltime int64) int64
	LoadSubsByID(id primitive.ObjectID) (*models.Subs, error)
	SaveSubs(subs *models.Subs) (*models.Subs, error)
	DeleteSubs(subs *models.Subs) (bool, error)
	FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error)
	FindSubsBetweenTime(fromTime int64, toTime int64) (*[]models.Subs, error)
	FindSubsExist(subs *models.Subs) (*models.Subs, error)
}

type SubsServiceImpl struct {
	subsRepository repository.SubsRepository
	lastTime       int64
}

func NewSubsService(subsRepository repository.SubsRepository) SubsService {
	return &SubsServiceImpl{
		subsRepository: subsRepository,
	}
}

func (s SubsServiceImpl) CreateSubs(chatID int64, cTime int64, uid primitive.ObjectID, latitude float64, longitude float64, frequency int64) models.Subs {
	if cTime == 0 {
		cTime = s.GetUTCTimeStamp64()
	}
	subs := s.subsRepository.CreateSubs()
	subs.ID = primitive.NewObjectID()
	subs.UID = uid
	subs.ChatID = chatID
	subs.CTime = cTime
	subs.UTime = cTime
	subs.LTime = 0
	subs.NTime = cTime + frequency
	subs.Latitude = latitude
	subs.Longitude = longitude
	subs.Frequency = frequency
	return subs
}

func (s SubsServiceImpl) GetUTCTimeStamp64() int64 {
	return time.Now().UTC().Unix()
}

func (s SubsServiceImpl) SetLastTime(ltime int64) int64 {
	if ltime == 0 {
		s.lastTime = s.GetUTCTimeStamp64()
	} else {
		s.lastTime = ltime
	}
	return s.lastTime
}

func (s SubsServiceImpl) GetLastTime() int64 {
	return s.lastTime
}

func (s SubsServiceImpl) SaveSubs(subs *models.Subs) (*models.Subs, error) {
	ss, err := s.subsRepository.FindUserExistSubs(subs)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription save error")
	}
	if ss == nil {
		ss, err = s.subsRepository.AddSubs(subs)
	} else {
		ss, err = s.subsRepository.UpdateSubs(subs)
	}
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription save error")
	}
	return ss, nil
}

func (s SubsServiceImpl) LoadSubsByID(id primitive.ObjectID) (*models.Subs, error) {
	sl, err := s.subsRepository.LoadSubs(id)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription load error")
	}
	return sl, nil
}

func (s SubsServiceImpl) DeleteSubs(subs *models.Subs) (bool, error) {
	err := s.subsRepository.DeleteSubs(subs)
	if err != nil {
		logrus.Debug(err)
		return false, fmt.Errorf("subscription delete error")
	}
	return true, nil
}

func (s SubsServiceImpl) FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error) {
	su, err := s.subsRepository.FindSubsByUID(uid)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription find by uid error")
	}
	return su, nil
}

func (s SubsServiceImpl) UserHasSubscribe(uid primitive.ObjectID, latitude float64, longitude float64) (*[]models.Subs, error) {
	su, err := s.subsRepository.FindSubsByUID(uid)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription find by uid error")
	}
	return su, nil
}

func (s SubsServiceImpl) FindSubsBetweenTime(fromTime int64, toTime int64) (*[]models.Subs, error) {
	su, err := s.subsRepository.FindNextSubsByTime(fromTime, toTime)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("subscription find between time error")
	}
	return su, nil
}

func (s SubsServiceImpl) FindSubsExist(subs *models.Subs) (*models.Subs, error) {
	ss, err := s.subsRepository.FindUserExistSubs(subs)
	if err != nil {
		logrus.Debug(err)
		return nil, fmt.Errorf("find exist subscription error")
	}
	return ss, nil
}
