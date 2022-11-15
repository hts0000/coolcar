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
	"math/rand"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()

	pm := &profileManager{}
	cm := &carManager{}
	s := newService(c, t, pm, cm)

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  312.12,
			Longitude: 123.123,
		},
	}
	pm.iID = "identity1"
	golden := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":312.12,"longitude":123.123},"poi_name":"迪士尼","timestamp_sec":1666017762},"current":{"location":{"latitude":312.12,"longitude":123.123},"poi_name":"迪士尼","timestamp_sec":1666017762},"status":1,"identity_id":"identity1"}`
	nowFunc = func() int64 {
		return 1666017762
	}
	cases := []struct {
		name         string
		aid          id.AccountID
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:    "normal_create",
			aid:     "account1",
			tripID:  "632b1c6e130f50c2748137aa",
			want:    fmt.Sprintf(golden, "account1"),
			wantErr: false,
		},
		{
			name:       "profile_error",
			aid:        "account2",
			tripID:     "632b1c6e130f50c2748137ab",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_error",
			aid:          "account3",
			tripID:       "632b1c6e130f50c2748137ac",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_error",
			aid:          "account4",
			tripID:       "632b1c6e130f50c2748137ad",
			carUnlockErr: fmt.Errorf("unlock"),
			want:         fmt.Sprintf(golden, "account4"),
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
			c := auth.ContextWithAccountID(context.Background(), cc.aid)
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
	pm := &profileManager{}
	cm := &carManager{}
	s := newService(context.Background(), t, pm, cm)

	cases := []struct {
		name    string
		aid     id.AccountID
		tripID  id.TripID
		wantErr bool
	}{
		{
			name: "normal_get",
			aid:  "account1",
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
			tr, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
				CarId: "car_1",
				Start: &rentalpb.Location{},
			})
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
	pm := &profileManager{}
	cm := &carManager{}
	s := newService(context.Background(), t, pm, cm)

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

func TestTripLifecycle(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("account_for_lifecycle"))
	s := newService(c, t, &profileManager{}, &carManager{})

	tid := id.TripID("632b1c6e140f50c2748137ad")
	mgutil.NewObjectIDWithValue(tid)
	cases := []struct {
		name string
		now  int64
		op   func() (*rentalpb.Trip, error)
		want *rentalpb.Trip
	}{
		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					CarId: "car1",
					Start: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: &rentalpb.Trip{
				AccountId: "account_for_lifecycle",
				CarId:     "car1",
				Start: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
					PoiName:      "天安门",
					TimestampSec: 10000,
				},
				Current: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
					PoiName:      "天安门",
					TimestampSec: 10000,
				},
				Status: rentalpb.TripStatus_IN_PROGRESS,
			},
			// `{"account_id":"account_for_lifecycle","car_id":"car1",
			// "start":{"location":{"latitude":22.213,"longitude":33.212},"poi_name":"天安门","timestamp_sec":10000},
			// "current":{"location":{"latitude":22.213,"longitude":33.212},"poi_name":"天安门","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
				})
			},
			want: &rentalpb.Trip{
				AccountId: "account_for_lifecycle",
				CarId:     "car1",
				Start: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
					PoiName:      "天安门",
					TimestampSec: 10000,
				},
				Current: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
					FeeCent:      6828,
					KmDriven:     262.73474627535245,
					PoiName:      "中关村",
					TimestampSec: 20000,
				},
				Status: rentalpb.TripStatus_IN_PROGRESS,
			},
			// `{"account_id":"account_for_lifecycle","car_id":"car1",
			// "start":{"location":{"latitude":22.213,"longitude":33.212},"poi_name":"天安门","timestamp_sec":10000},
			// "current":{"location":{"latitude":28.123,"longitude":312.123},"fee_cent":6828,"km_driven":262.73474627535245,"poi_name":"中关村","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: &rentalpb.Trip{
				AccountId: "account_for_lifecycle",
				CarId:     "car1",
				Start: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
					PoiName:      "天安门",
					TimestampSec: 10000,
				},
				Current: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
					FeeCent:      6929,
					KmDriven:     340.34172200547056,
					PoiName:      "中关村",
					TimestampSec: 30000,
				},
				End: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
					FeeCent:      6929,
					KmDriven:     340.34172200547056,
					PoiName:      "中关村",
					TimestampSec: 30000,
				},
				Status: rentalpb.TripStatus_FINISHED,
			},
			//`{"account_id":"account_for_lifecycle","car_id":"car1",
			// "start":{"location":{"latitude":22.213,"longitude":33.212},"poi_name":"天安门","timestamp_sec":10000},
			// "current":{"location":{"latitude":28.123,"longitude":312.123},"fee_cent":6929,"km_driven":340.34172200547056,"poi_name":"中关村","timestamp_sec":30000},
			// "end":{"location":{"latitude":28.123,"longitude":312.123},"fee_cent":6929,"km_driven":340.34172200547056,"poi_name":"中关村","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: &rentalpb.Trip{
				AccountId: "account_for_lifecycle",
				CarId:     "car1",
				Start: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  22.213,
						Longitude: 33.212,
					},
					PoiName:      "天安门",
					TimestampSec: 10000,
				},
				Current: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
					FeeCent:      6929,
					KmDriven:     340.34172200547056,
					PoiName:      "中关村",
					TimestampSec: 30000,
				},
				End: &rentalpb.LocationStatus{
					Location: &rentalpb.Location{
						Latitude:  28.123,
						Longitude: 312.123,
					},
					FeeCent:      6929,
					KmDriven:     340.34172200547056,
					PoiName:      "中关村",
					TimestampSec: 30000,
				},
				Status: rentalpb.TripStatus_FINISHED,
			},
			// `{"account_id":"account_for_lifecycle","car_id":"car1",
			// "start":{"location":{"latitude":22.213,"longitude":33.212},"poi_name":"天安门","timestamp_sec":10000},
			// "current":{"location":{"latitude":28.123,"longitude":312.123},"fee_cent":6929,"km_driven":340.34172200547056,"poi_name":"中关村","timestamp_sec":30000},
			// "end":{"location":{"latitude":28.123,"longitude":312.123},"fee_cent":6929,"km_driven":340.34172200547056,"poi_name":"中关村","timestamp_sec":30000},"status":2}`,
		},
	}
	rand.Seed(1314)
	for _, cc := range cases {
		nowFunc = func() int64 {
			return cc.now
		}
		got, err := cc.op()
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
			continue
		}
		if diff := cmp.Diff(got, cc.want, protocmp.Transform()); diff != "" {
			// -号的行是期望得到的
			// +号的行是得到的行
			t.Errorf("result differs: -want +got: %s", diff)
		}
	}
}

func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
	mc, err := mongotesting.NewClient(context.Background())
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}
	db := mc.Database("coolcar")
	mongotesting.SetupIndexs(c, db)
	return &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(db),
		Logger:         logger,
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
	lockErr   error
}

func (m *carManager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return m.verifyErr
}

func (m *carManager) Unlock(context.Context, id.CarID, id.AccountID, id.TripID, string) error {
	return m.unlockErr
}

func (m *carManager) Lock(context.Context, id.CarID) error {
	return m.lockErr
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m,
		mongotesting.URIConfig{
			IP:   "182.61.47.223",
			Port: "9876",
		},
	))
}
