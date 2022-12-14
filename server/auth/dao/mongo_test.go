package dao

import (
	"context"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m,
		mongotesting.URIConfig{
			IP:   "182.61.47.223",
			Port: "9876",
		},
	))
}

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	_, err = m.col.InsertMany(c, []any{
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.AccountID("632b1c6e130f50c2748137aa")),
			openIDField:        "open_id_1",
		},
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.AccountID("632b1c6e130f50c2748137ab")),
			openIDField:        "open_id_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}
	mgutil.NewObjectIDWithValue(id.AccountID("632b1c6e130f50c2748137ad"))

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
			if id.String() != c.want {
				t.Errorf("faild resolve account id: want: %q, got: %q", c.want, id)
			}
		})
	}
}
