package main

import (
	"RiYu/mongodb"
	"RiYu/server/def"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 查共有几种『标』
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
	project := bson.D{{"$project", bson.M{"_id": 0, "标": 1}}}
	unwind := bson.D{{"$unwind", "$标"}}
	group := bson.D{{"$group", bson.D{{"_id", "$标"}}}}
	cursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{project, unwind, group})
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%s\n", result["_id"])
	}
}
