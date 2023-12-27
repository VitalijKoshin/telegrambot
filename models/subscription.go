package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subs struct {
	ID        primitive.ObjectID `bson:"_id"`
	UID       primitive.ObjectID `bson:"user_id"`
	ChatID    int64              `bson:"chat_id"`
	Latitude  float64            `bson:"geo_latitude"`
	Longitude float64            `bson:"geo_longitude"`
	CTime     int64              `bson:"created"`
	UTime     int64              `bson:"updated"`
	LTime     int64              `bson:"last_time"`
	NTime     int64              `bson:"next_time"`
	Frequency int64              `bson:"frequency"`
}
