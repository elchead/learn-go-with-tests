package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// err = client.Ping(ctx, readpref.Primary())
	// assert.Equal(t, nil, err)

	name := "Kate"
	store := NewMongoStore(ctx, "mongodb://localhost:27017")
	err := store.PostPlayerWin(name)
	assert.NoError(t, err)
	score, ok := store.GetPlayerScore(name)
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, score)
}
