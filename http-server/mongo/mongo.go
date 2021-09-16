package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name  string
	Score int
}

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoStore(ctx context.Context, uri string) *MongoStore {
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection := cl.Database("tdd").Collection("players")
	return &MongoStore{cl, collection, ctx}
}

func (s *MongoStore) GetPlayerScore(name string) (int, bool) {
	doc := s.collection.FindOne(s.ctx, bson.M{"name": name})
	readPlayer := &Player{}
	if err := doc.Decode(readPlayer); err != nil {
		return 0, false
	}
	return readPlayer.Score, true
}

func (s *MongoStore) PostPlayerWin(name string) error {
	player := Player{Name: name, Score: 1}
	_, err := s.collection.InsertOne(s.ctx, player)
	return err
}
