package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Player struct {
	Name  string
	Score int
}

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	assert.Equal(t, nil, err)
	collection := client.Database("tdd").Collection("players")

	player := Player{Name: "Bob"}
	collection.InsertOne(ctx, player)
	doc := collection.FindOne(ctx, bson.D{})

	readPlayer := &Player{}
	err = doc.Decode(readPlayer)
	assert.NoError(t, err)
	assert.Equal(t, "Bob", readPlayer.Name)
}
