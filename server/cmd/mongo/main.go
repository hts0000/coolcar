package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://182.61.47.223:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mc.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	col := mc.Database("coolcar").Collection("coolcar")
	query(c, col)
}

func query(c context.Context, col *mongo.Collection) {
	cur, err := col.Find(c, bson.D{})
	if err != nil {
		panic(err)
	}
	for cur.Next(c) {
		var res struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenID string             `bson:"open_id"`
		}
		cur.Decode(&res)
		fmt.Printf("%+v\n", res)
	}
}

func insert(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(c, []any{
		bson.M{"open_id": "1"},
		bson.M{"open_id": "2"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
