package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField    = "trip"
	accountField = tripField + ".accountid"
	statusField  = tripField + ".status"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection(tripField),
	}
}

type TripRecord struct {
	mgutil.IDField        `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgutil.NewObjID()
	r.UpdateAt = mgutil.UpdatedAt()

	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *Mongo) GetTrip(c context.Context, id id.TripID, accountID id.AccountID) (*TripRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	res := m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
		accountField:       accountID,
	})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %v", err)
	}

	return &tr, nil
}

func (m *Mongo) GetTrips(c context.Context, aid id.AccountID, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountField: aid.String(),
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}

	res, err := m.col.Find(c, filter)
	if err != nil {
		return nil, err
	}

	var trips []*TripRecord
	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}

func (m *Mongo) UpdateTrip(c context.Context, tid id.TripID, aid id.AccountID, updatedAt int64, trip *rentalpb.Trip) error {
	objID, err := objid.FromID(tid)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	newUpdatedAt := mgutil.UpdatedAt()
	res, err := m.col.UpdateOne(c, bson.M{
		mgutil.IDFieldName: objID,
		accountField:       aid.String(),
		// ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
		// ?????????????????????
		mgutil.UpdatedAtFieldName: updatedAt,
	}, mgutil.Set(bson.M{
		tripField: trip,
		// ?????????????????????????????????????????????????????????????????????????????????
		mgutil.UpdatedAtFieldName: newUpdatedAt,
	}))
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
