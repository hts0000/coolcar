package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateTrip(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("account1"))
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	pm := profileManager{}
	cm := carManager{}
	s := &Service{
		ProfileManager: &pm,
		CarManager:     &cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  312.12,
			Longitude: 123.123,
		},
	}
	pm.iID = "identity1"
	golden := `{"account_id":"account1","car_id":"car1","start":{"location":{"latitude":312.12,"longitude":123.123},"poi_name":"迪士尼"},"current":{"location":{"latitude":312.12,"longitude":123.123},"poi_name":"迪士尼"},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:    "normal_create",
			tripID:  "632b1c6e130f50c2748137aa",
			want:    golden,
			wantErr: false,
		},
		{
			name:       "profile_error",
			tripID:     "632b1c6e130f50c2748137ab",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_error",
			tripID:       "632b1c6e130f50c2748137ac",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_error",
			tripID:       "632b1c6e130f50c2748137ad",
			carUnlockErr: fmt.Errorf("unlock"),
			want:         golden,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("%s: want error; got none", cc.name)
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("%s: cannot create trip: %v", cc.name, err)
				return
			}
			if res.Id != cc.tripID {
				t.Errorf("%s: incorrect id; want %q, got %q", cc.name, cc.tripID, res.Id)
			}
			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("%s: cannot marshal trip: %v", cc.name, err)
			}
			got := string(b)
			if cc.want != got {
				t.Errorf("%s: incorrect response: want %s, got %s", cc.name, cc.want, got)
			}
		})
	}
}

func TestGetTrip(t *testing.T) {
	mc, err := mongotesting.NewClient(context.Background())
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	pm := profileManager{}
	cm := carManager{}
	s := &Service{
		ProfileManager: &pm,
		CarManager:     &cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}

	cases := []struct {
		name    string
		aid     id.AccountID
		tripID  id.TripID
		wantErr bool
	}{
		{
			name: "normal_get",
			aid:  "account2",
		},
		{
			name:    "error_get",
			aid:     "account2",
			tripID:  id.TripID("NON-EXIST_TRIP_ID"),
			wantErr: true,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			c := auth.ContextWithAccountID(context.Background(), cc.aid)
			tr, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{CarId: "car_1"})
			if err != nil {
				t.Fatalf("%s: cannot create trip: %v", cc.name, err)
			}
			req := &rentalpb.GetTripRequest{
				Id: tr.Id,
			}
			if cc.tripID != "" {
				req.Id = cc.tripID.String()
			}
			trip, err := s.GetTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Fatalf("%s: want error; got none", cc.name)
				}
				return
			}
			if err != nil {
				t.Fatalf("%s: cannot get trip: %v", cc.name, err)
			}
			if diff := cmp.Diff(tr.Trip, trip, protocmp.Transform()); diff != "" {
				// -号的行是期望得到的
				// +号的行是得到的行
				t.Errorf("result differs: -want +got: %s", diff)
			}
		})
	}
}

func TestGetTrips(t *testing.T) {
	mc, err := mongotesting.NewClient(context.Background())
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	pm := profileManager{}
	cm := carManager{}
	s := &Service{
		ProfileManager: &pm,
		CarManager:     &cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}

	trips := []*rentalpb.Trip{
		{
			AccountId: "account3",
			CarId:     "car_1",
			Status:    rentalpb.TripStatus_FINISHED,
		},
		{
			AccountId: "account3",
			CarId:     "car_2",
			Status:    rentalpb.TripStatus_FINISHED,
		},
		{
			AccountId: "account3",
			CarId:     "car_3",
			Status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	for _, trip := range trips {
		c := auth.ContextWithAccountID(context.Background(), id.AccountID(trip.AccountId))
		_, err := s.Mongo.CreateTrip(c, trip)
		if err != nil {
			t.Fatalf("cannot create trip: %v", err)
		}
	}

	cases := []struct {
		name    string
		aid     id.AccountID
		status  rentalpb.TripStatus
		wantCnt int
	}{
		{
			name:    "get_finished",
			aid:     "account3",
			status:  rentalpb.TripStatus_FINISHED,
			wantCnt: 2,
		},
		{
			name:    "get_in_progress",
			aid:     "account3",
			status:  rentalpb.TripStatus_IN_PROGRESS,
			wantCnt: 1,
		},
		{
			name:    "get_all",
			aid:     "account3",
			status:  rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCnt: 3,
		},
		{
			name:    "get_none",
			aid:     "NON-EXIST_ACCOUNT_ID",
			status:  rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCnt: 0,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			c := auth.ContextWithAccountID(context.Background(), id.AccountID(cc.aid))
			resp, err := s.GetTrips(c, &rentalpb.GetTripsRequest{
				Status: cc.status,
			})
			if err != nil {
				t.Fatalf("%s: cannot get trips: %v", cc.name, err)
			}
			if len(resp.Trips) != cc.wantCnt {
				t.Fatalf("%s: get trips want cnt %d, got %d", cc.name, cc.wantCnt, len(resp.Trips))
			}
		})
	}
}

// 提供一个实现者，直接返回测试数据，固化流程中的随机值
type profileManager struct {
	iID id.IdentityID
	err error
}

func (p *profileManager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return p.iID, p.err
}

type carManager struct {
	verifyErr error
	unlockErr error
}

func (c *carManager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return c.verifyErr
}

func (c *carManager) Unlock(context.Context, id.CarID) error {
	return c.unlockErr
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m,
		mongotesting.URIConfig{
			IP:   "182.61.47.223",
			Port: "9876",
		},
	))
}
