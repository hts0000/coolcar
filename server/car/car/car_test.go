package car

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/dao"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"coolcar/shared/server"
	"encoding/json"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, mongotesting.URIConfig{
		IP:   "182.61.47.223",
		Port: "9876",
	}))
}

func TestCarUpdate(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	s := &Service{
		Logger: logger,
		Mongo:  dao.NewMongo(mc.Database("coolcar")),
	}

	carID := id.CarID("632b1c6e130f50c2748137ab")
	mgutil.NewObjectIDWithValue(carID)
	_, err = s.CreateCar(c, &carpb.CreateCarRequest{})
	if err != nil {
		t.Fatalf("cannot create car: %v", err)
	}

	cases := []struct {
		name    string
		op      func() error
		want    string
		wantErr bool
	}{
		{
			name: "get_car",
			op: func() error {
				return nil
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unlock_car",
			op: func() error {
				_, err := s.UnlockCar(c, &carpb.UnlockCarRequest{
					Id:     carID.String(),
					TripId: "test_trip",
					Driver: &carpb.Driver{
						Id:        "test_driver",
						AvatarUrl: "test_avatar",
					},
				})
				return err
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unlock_complete",
			op: func() error {
				_, err := s.UpdateCar(c, &carpb.UpdateCarRequest{
					Id: carID.String(),
					Position: &carpb.Location{
						Latitude:  31,
						Longitude: 121,
					},
					Status: carpb.CarStatus_UNLOCKED,
				})
				return err
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unlock_car_by_another_driver",
			op: func() error {
				_, err := s.UnlockCar(c, &carpb.UnlockCarRequest{
					Id:     carID.String(),
					TripId: "bad_trip",
					Driver: &carpb.Driver{
						Id:        "bad_driver",
						AvatarUrl: "test_avatar",
					},
				})
				return err
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "lock_car",
			op: func() error {
				_, err := s.LockCar(c, &carpb.LockCarRequest{
					Id: carID.String(),
				})
				return err
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "complete_lock_car",
			op: func() error {
				_, err := s.UpdateCar(c, &carpb.UpdateCarRequest{
					Id:     carID.String(),
					Status: carpb.CarStatus_LOCKED,
				})
				return err
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, cc := range cases {
		err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want err; got none", cc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
			continue
		}
		car, err := s.GetCar(c, &carpb.GetCarRequest{
			Id: carID.String(),
		})
		if err != nil {
			t.Errorf("%s: cannot get car after operation: %v", cc.name, err)
		}
		b, err := json.Marshal(car)
		if err != nil {
			t.Errorf("%s: failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s: incorrect response; want: %s, got: %s", cc.name, cc.want, got)
		}
	}
}
