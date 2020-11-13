package service

import (
	dtype "RiYu/data/type"
	"RiYu/mongodb"
	"RiYu/server/def"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

var (
	collection *mongo.Collection
	verbs      []*dtype.Words
	adjs       []*dtype.Words
	advs       []*dtype.Words
	others     []*dtype.Words
	words      [][]*dtype.Words
)

// Init 准备
func Init(client *mongodb.Client) error {
	collection = client.Database(def.WordsDB).Collection(def.WordsCol)
	// 将词库读入内存
	cursor, err := collection.Find(context.Background(), bson.M{"标": bson.M{"$in": []string{"五段", "一段", "サ变", "カ变"}}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.Background(), &verbs); err != nil {
		return err
	}
	if cursor, err = collection.Find(context.Background(), bson.M{"标": bson.M{"$regex": "(?:形|形动)"}}); err != nil {
		return err
	}
	if err = cursor.All(context.Background(), &adjs); err != nil {
		return err
	}
	if cursor, err = collection.Find(context.Background(), bson.M{"标": "副"}); err != nil {
		return err
	}
	if err = cursor.All(context.Background(), &advs); err != nil {
		return err
	}
	if cursor, err = collection.Find(context.Background(), bson.M{"标": nil}); err != nil {
		return err
	}
	if err = cursor.All(context.Background(), &others); err != nil {
		return err
	}
	words = [][]*dtype.Words{verbs, adjs, advs, others}
	return nil
}

// Router 路由
func Router() *httprouter.Router {
	router := httprouter.New()
	for _, data := range routers {
		router.Handle(data.Method, data.Path, data.Handler)
	}
	return router
}
