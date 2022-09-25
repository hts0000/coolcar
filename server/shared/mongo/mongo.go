package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const IDField = "_id"

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

func Set(v any) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v any) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
