package main

import (
	"RiYu/mongodb"
	"RiYu/server/def"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := mongodb.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Stop(); err != nil {
			panic(err)
		}
	}()
	collection := client.Database(def.WordsDB).Collection(def.WordsCol)
	dupe(collection, "日")
	dupe(collection, "汉")
}

func dupe(collection *mongo.Collection, key string) {
	prop := "$" + key
	project := bson.D{{"$project", bson.M{"_id": 0, key: bson.D{{"$split", []string{prop, "・"}}}}}}
	unwind := bson.D{{"$unwind", prop}}
	group := bson.D{{"$group", bson.D{{"_id", prop}, {"total", bson.M{"$sum": 1}}}}}
	match := bson.D{{"$match", bson.M{"total": bson.M{"$gt": 1}}}}
	cursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{project, unwind, group, match})
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}
	fmt.Printf("%s：\n", key)
	for _, result := range results {
		fmt.Printf("%s\n", result["_id"])
	}
}
