package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	uri := "mongodb://localhost:27017/?retryWrites=true&writeConcern=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Disconnected")
	}()

	db := client.Database("sample_guides")
	coll := db.Collection("comets")

	// insert code goes here
	docs := []interface{}{
		bson.D{{"name", "Halley's Comet"},
			{"officialName", "1P/Halley"},
			{"orbitalPeriod", 75},
			{"radius", 3.4175},
			{"mass", 2.2e14}},
		bson.D{{"name", "Wild2"},
			{"officialName", "81P/Wild"},
			{"orbitalPeriod", 6.41},
			{"radius", 1.5534},
			{"mass", 2.3e13}},
		bson.D{{"name", "Comet Hyakutake"},
			{"officialName", "C/1996 B2"},
			{"orbitalPeriod", 17000},
			{"radius", 0.77671},
			{"mass", 8.8e12}},
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
}
