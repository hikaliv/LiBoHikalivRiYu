package main

import (
	"RiYu/data/static"
	dtype "RiYu/data/type"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	flag.Parse()
	client, err := static.ConnectDB()
	if err != nil {
		panic(err)
	}
	// 断联客户端
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	// 命令行为
	// go run save.go xxx.json yyy.json
	// 读所有 json 文件，或
	// go run save.go zzz
	// 读 zzz 下所有 json 文件
	static.LookJSON(flag.Args(), func(filename string) error {
		return save(filename, client.Database("Japanese").Collection("Words"))
	})
}

func save(filename string, collection *mongo.Collection) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("文件 %s 读时错：%v", filename, err)
	}
	var relations []*dtype.Relations
	if err = json.Unmarshal(data, &relations); err != nil {
		return fmt.Errorf("文件 %s 解析数据错误：%v", filename, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var (
		result     *mongo.UpdateResult
		unMatched  []string
		unModified []string
		updated    []string
	)
	for _, relation := range relations {
		var (
			needtoUpdate = false
			filter       = bson.M{"id": relation.ID}
			existedDoc   bson.M
		)
		if err = collection.FindOne(ctx, filter).Decode(&existedDoc); err != nil {
			return err
		} else if lian, ok := existedDoc["联"]; !ok {
			needtoUpdate = true
		} else if exRelMap, ok := lian.(bson.M); !ok {
			return fmt.Errorf("处理文件 %s 时解析库中 %s 键发现值无法读取", filename, relation.ID)
		} else {
			for key, value := range relation.Relation {
				if existedValue, ok := exRelMap[key]; !ok {
					needtoUpdate = true
					break
				} else if existedValue != value {
					needtoUpdate = true
					break
				}
			}
		}
		if !needtoUpdate {
			unModified = append(unModified, relation.ID)
			continue
		}
		update := bson.M{"$set": bson.M{"联": relation.Relation}}
		if result, err = collection.UpdateOne(ctx, filter, update); err != nil {
			return err
		}
		if result.MatchedCount == 0 {
			unMatched = append(unMatched, relation.ID)
		} else if result.ModifiedCount == 0 {
			unModified = append(unModified, relation.ID)
		} else {
			updated = append(updated, relation.ID)
		}
	}
	fmt.Printf("待检查者 %d 个\n", len(relations))
	if unMatched != nil {
		fmt.Printf("在库里不存在的标识共有 %d 个：\n%v\n", len(unMatched), unMatched)
	}
	if unModified != nil {
		fmt.Printf("内容未更新的标识有 %d 个：\n%v\n", len(unModified), unModified)
	}
	if updated != nil {
		fmt.Printf("更新的标识有 %d 个：\n%v\n", len(updated), updated)
	}
	return nil
}
