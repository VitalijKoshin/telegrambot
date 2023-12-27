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

type SubsRepositoryImpl struct {
	subsCollection *mongo.Collection
	ctx            context.Context
}

func NewSubsRepository(cfg *config.Config, mdbClient mongodb.MongoClient) repository.SubsRepository {
	collection := mdbClient.GetClient().Database(cfg.MongoDbName).Collection(repository.NameSubsCollection)
	return &SubsRepositoryImpl{
		subsCollection: collection,
		ctx:            context.TODO(),
	}
}

func (s *SubsRepositoryImpl) CreateSubs() models.Subs {
	return models.Subs{}
}

func (s *SubsRepositoryImpl) AddSubs(subs *models.Subs) (*models.Subs, error) {
	if subs.ID.IsZero() {
		subs.ID = primitive.NewObjectID()
	}
	_, err := s.subsCollection.InsertOne(s.ctx, subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (s *SubsRepositoryImpl) LoadSubs(ID primitive.ObjectID) (*models.Subs, error) {
	var subsFound models.Subs
	filter := bson.M{"_id": ID}
	err := s.subsCollection.FindOne(s.ctx, filter).Decode(&subsFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &subsFound, nil
}

func (s *SubsRepositoryImpl) UpdateSubs(subs *models.Subs) (*models.Subs, error) {
	filter := bson.M{"_id": subs.ID}
	uByte, err := bson.Marshal(subs)
	if err != nil {
		return nil, err
	}
	var update bson.M
	err = bson.Unmarshal(uByte, &update)
	if err != nil {
		return nil, err
	}
	_, err = s.subsCollection.UpdateOne(s.ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (s *SubsRepositoryImpl) DeleteSubs(subs *models.Subs) error {
	_, err := s.subsCollection.DeleteOne(s.ctx, subs)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubsRepositoryImpl) FindSubsByUID(uid primitive.ObjectID) (*[]models.Subs, error) {
	var subsFound []models.Subs
	filter := bson.D{{Key: "user_id", Value: uid}}
	cursor, err := s.subsCollection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.ctx, &subsFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &subsFound, nil
		}
		return nil, err
	}
	return &subsFound, nil
}

func (s *SubsRepositoryImpl) FindSubsByUIDLatiLongi(uid primitive.ObjectID, latitude float64, longitude float64) (*[]models.Subs, error) {
	var subsFound []models.Subs
	filter := bson.D{{Key: "user_id", Value: uid}}
	cursor, err := s.subsCollection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.ctx, &subsFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &subsFound, nil
		}
		return nil, err
	}
	return &subsFound, nil
}

func (s *SubsRepositoryImpl) FindUserExistSubs(subs *models.Subs) (*models.Subs, error) {
	var subsFound models.Subs
	filter := bson.D{{Key: "user_id", Value: subs.UID}, {Key: "geo_latitude", Value: subs.Latitude}, {Key: "geo_longitude", Value: subs.Longitude}}
	err := s.subsCollection.FindOne(s.ctx, filter).Decode(&subsFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &subsFound, nil
}

func (s *SubsRepositoryImpl) FindNextSubsByTime(fromTime int64, toTime int64) (*[]models.Subs, error) {
	var subsFound []models.Subs
	filter := bson.D{{Key: "next_time", Value: bson.D{{Key: "$lt", Value: toTime}, {Key: "$gte", Value: fromTime}}}}
	cursor, err := s.subsCollection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.ctx, &subsFound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &subsFound, nil
		}
		return nil, err
	}
	return &subsFound, nil
}
