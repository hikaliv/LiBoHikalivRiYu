package main

// go run save.go xxx.json yyy.json
// go run save.go zzz/

import (
	"RiYu/data/static"
	dtype "RiYu/data/type"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	collection := client.Database("Japanese").Collection("Words")
	indexes := collection.Indexes()
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index().SetUnique(true),
	}
	// 对表建索引，当表不存在时，新建该表，当索引已存在时，若索引模式相同则无谓，否则冲突报错
	if _, err = indexes.CreateOne(context.Background(), indexModel, options.CreateIndexes().SetMaxTime(time.Second)); err != nil {
		panic(err)
	}
	// 其实可以利用这种闭包设计实现外部直接介入的数据统计
	total := 0
	// 命令行为
	// go run save.go xxx.json yyy.json
	// 读所有 json 文件，或
	// go run save.go zzz
	// 读 zzz 下所有 json 文件
	if err = static.LookJSON(flag.Args(), func(filename string) error {
		count, saveErr := save(filename, collection)
		total += count
		return saveErr
	}); err != nil {
		panic(err)
	}
	fmt.Printf("总共插入了 %d 条数据\n", total)
}

func save(filename string, collection *mongo.Collection) (count int, retErr error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		retErr = fmt.Errorf("文件 %s 读时错：%v", filename, err)
		return
	}
	var words []*dtype.Words
	if err = json.Unmarshal(data, &words); err != nil {
		retErr = fmt.Errorf("文件 %s 解析数据错误：%v", filename, err)
		return
	}
	// go mongodb 的 InsertMany 需要 []interface{} 型
	interfaces := make([]interface{}, len(words))
	// 生成以『汉日』两字段合并的散列值为 ID
	for itor := range words {
		word := words[itor]
		var buf bytes.Buffer
		buf.WriteString(word.Han)
		buf.WriteString(word.Ri)
		hash := sha256.Sum256(buf.Bytes())
		word.ID = base58.Encode(hash[:])
		interfaces[itor] = word
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertMany(ctx, interfaces, options.InsertMany().SetOrdered(false))
	// 判断错误是不是重复插入，如果是重复插入则忽略，其它抛出
	if err != nil {
		bulkErr, ok := err.(mongo.BulkWriteException)
		if !ok {
			retErr = fmt.Errorf("文件 %s 数据存时错：%v", filename, err)
			return
		}
		for _, writeError := range bulkErr.WriteErrors {
			if writeError.Code != 11000 {
				retErr = fmt.Errorf("文件 %s 数据存时错：%v", filename, err)
				return
			}
		}
	}
	count = len(result.InsertedIDs)
	fmt.Printf("文件 %s 插入了 %d 条数据\n", filename, count)
	return
}
