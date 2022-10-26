package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/blob"
	"coolcar/blob/cos"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://182.61.47.223:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	db := mongoClient.Database("coolcar")

	st, err := cos.NewService(
		"https://coolcar-1300912551.cos.ap-guangzhou.myqcloud.com",
		"https://cos.ap-guangzhou.myqcloud.com",
		"AKIDrdAUXKq69xVqwlV1HH0RguxlPpz50kHc",
		"B6kALI5c9QSDsLdiZCYdXBanmW6TbS3R",
	)
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   ":8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Mongo:   dao.NewMongo(db),
				Logger:  logger,
				Storage: st,
			})
		},
	}))
}
