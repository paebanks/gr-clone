package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client *mongo.Client
	ctx    context.Context
}

func GetClient() *mongo.Client {
	// Replace the uri string with your MongoDB deployment's connection string.
	/*
		uri := "mongodb://localhost:27017/gb-clone"
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
		// Ping the primary
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected and pinged.")

		var conn = MongoConnection{client: client, ctx: ctx}
		return conn
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return client
}

//var client = getClient();
