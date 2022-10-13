package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m,
		mongotesting.URIConfig{
			IP:   "182.61.47.223",
			Port: "9876",
		},
	))
}

func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	tid := id.TripID("6342bc318e2e86ef286e08b9")
	aid := id.AccountID("account_for_update")

	var now int64 = 10000
	mgutil.NewObjectIDWithValue(tid)
	mgutil.UpdatedAt = func() int64 {
		return now
	}

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	if tr.UpdateAt != 10000 {
		t.Fatalf("wrong updatedAt; want: 10000, got: %d", tr.UpdateAt)
	}

	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi_update",
		},
	}
	cases := []struct {
		name          string
		now           int64
		withUpdatedAt int64
		wantErr       bool
	}{
		{
			name:          "normal_update",
			now:           20000,
			withUpdatedAt: 10000,
		},
		{
			name:          "update_with_stale_timestamp",
			now:           30000,
			withUpdatedAt: 10000,
			wantErr:       true,
		},
		{
			name:          "update_with_refetch",
			now:           40000,
			withUpdatedAt: 20000,
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdatedAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		} else if err != nil {
			t.Errorf("%s: cannot update: %v", cc.name, err)
		}
		updatedTrip, err := m.GetTrip(c, tid, aid)
		if err != nil {
			t.Errorf("%s: cannot get trip after update: %v", cc.name, err)
		}
		if cc.now != updatedTrip.UpdateAt {
			t.Errorf("%s: incorrect updatedat; want: %d, got: %d", cc.name, cc.now, updatedTrip.UpdateAt)
		}
	}
}

// 测试有索引的情况下行程的建立
func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	db := mc.Database("coolcar")
	err = mongotesting.SetupIndexs(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexs: %v", err)
	}

	m := NewMongo(db)
	// 在建立了索引的db中根据accountID创建行程，
	// 成功了应该返回tripID，
	// 失败了应该得到error
	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			tripID:     "6342bc318e2e89ef286e08b9",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
			wantErr:    false,
		},
		{
			name:       "another_finished",
			tripID:     "6342bc318e2e89ef286e08b8",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
			wantErr:    false,
		},
		{
			name:       "in_progress",
			tripID:     "6342bc318e2e89ef286e08b7",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    false,
		},
		{
			name:       "another_in_progress",
			tripID:     "6342bc318e2e89ef286e08b6",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account",
			tripID:     "6342bc318e2e89ef286e08b5",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    false,
		},
	}

	for _, cc := range cases {
		mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: error creating trip: %v", cc.name, err)
			continue
		}
		if tr.ID.Hex() != cc.tripID {
			t.Errorf("%s: incorrect trip id; want: %q, got: %q", cc.name, cc.tripID, tr.ID.Hex())
		}
	}
}

func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	acct := id.AccountID("account1")
	mgutil.NewObjID = primitive.NewObjectID
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "start_point",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "end_point",
			FeeCent:  10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}

	got, err := m.GetTrip(c, objid.ToTripID(tr.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}

	if diff := cmp.Diff(tr, got, protocmp.Transform()); diff != "" {
		// -号的行是期望得到的
		// +号的行是得到的行
		t.Errorf("result differs: -want +got: %s", diff)
	}
}

func TestGetTrips(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))

	rows := []struct {
		id        string
		accountID string
		status    rentalpb.TripStatus
	}{
		{
			id:        "7342bc318e2e89ef286e08b9",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "7342bc318e2e89ef286e08b8",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "7342bc318e2e89ef286e08b7",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "7342bc318e2e89ef286e08b6",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "7342bc318e2e89ef286e08b5",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	for _, r := range rows {
		mgutil.NewObjectIDWithValue(id.TripID(r.id))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: r.accountID,
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCount  int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount: 4,
		},
		{
			name:       "get_in_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
			wantOnlyID: "7342bc318e2e89ef286e08b6",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(), id.AccountID(cc.accountID), cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}

			if len(res) != cc.wantCount {
				t.Errorf("%s: incorrect result count; want: %d, got: %d", cc.name, cc.wantCount, len(res))
			}
			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("%s: only_id incorrect; want: %q, got: %q", cc.name, cc.wantOnlyID, res[0].ID.Hex())
				}
			}
		})
	}
}
