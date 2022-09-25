package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m,
		mongotesting.URIConfig{
			IP:   "182.61.47.223",
			Port: "9876",
		},
		&mongoURI,
	))
}

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	_, err = m.col.InsertMany(c, []any{
		bson.M{
			mgo.IDField: mustObjID("632b1c6e130f50c2748137aa"),
			openIDField: "open_id_1",
		},
		bson.M{
			mgo.IDField: mustObjID("632b1c6e130f50c2748137ab"),
			openIDField: "open_id_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}
	m.newObjID = func() primitive.ObjectID {
		return mustObjID("632b1c6e130f50c2748137ad")
	}

	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "open_id_1",
			want:   "632b1c6e130f50c2748137aa",
		},
		{
			name:   "another_existing_user",
			openID: "open_id_2",
			want:   "632b1c6e130f50c2748137ab",
		},
		{
			name:   "new_user",
			openID: "open_id_3",
			want:   "632b1c6e130f50c2748137ad",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), c.openID)
			if err != nil {
				t.Errorf("faild resolve account id for %q: %v", c.openID, err)
			}
			if id != c.want {
				t.Errorf("faild resolve account id: want: %q, got: %q", c.want, id)
			}
		})
	}
}

func mustObjID(hex string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return id
}
