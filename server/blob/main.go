package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/blob"
	"coolcar/blob/cos"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"fmt"
	"log"
	"os"

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

	// 从环境变量中读取腾讯云认证id、key、存储桶名称和地区
	var (
		secID   = os.Getenv("SECRETID")
		secKey  = os.Getenv("SECRETKEY")
		bktName = os.Getenv("BUCKETNAME")
		region  = os.Getenv("REGION")

		bktAddr = fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bktName, region)
		serAddr = fmt.Sprintf("https://cos.%s.myqcloud.com", region)
	)

	st, err := cos.NewService(
		bktAddr,
		serAddr,
		secID,
		secKey,
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
