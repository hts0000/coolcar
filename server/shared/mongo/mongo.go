package mgutil

import (
	"coolcar/shared/mongo/objid"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdatedAtField struct {
	UpdateAt int64 `bson:"updatedat"`
}

var NewObjID = primitive.NewObjectID

func NewObjectIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
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
