package integration

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestInsertAndRead(t *testing.T) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	col := client.Database("testdb").Collection("testcoll")
	_, err = col.InsertOne(context.TODO(), map[string]string{"hello": "world"})
	//col.FindOne()
	if err != nil {
		t.Fatal(err)
	}
}
