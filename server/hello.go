package main

import (
	"coolcar/constraints"
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("sum base on generics:", sum([]float64{0.1, 0.2, 0.3}))

	// generate base on protoc: protoc -I="." --go_out="paths=source_relative:gen\go" .\trip.proto
	trip := trippb.Trip{
		Start:       "abc",
		End:         "cda",
		DurationSec: 1000,
		FeeCent:     1000,
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 115,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  110,
				Longitude: 110,
			},
			{
				Latitude:  100,
				Longitude: 100,
			},
		},
		Status: trippb.TripStatus_FINISHED,
	}
	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", b)
	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	b, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func sum[T constraints.Ordered](nums []T) (ans T) {
	for _, num := range nums {
		ans += num
	}
	return ans
}
