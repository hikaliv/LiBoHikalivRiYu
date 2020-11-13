package static

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDB 连数据库
func ConnectDB() (client *mongo.Client, err error) {
	// 定三秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// 多个 defer，先下后上
	defer cancel()
	// 启联客户端，试三秒
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		return
	}
	// 三秒之内，试联通与否
	err = client.Ping(ctx, readpref.Primary())
	return
}

// LookJSON 遍历文件目录
func LookJSON(args []string, save func(string) error) error {
	for _, filename := range args {
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			return fmt.Errorf("文件 %s 查属性时错：%v", filename, err)
		}
		if fileInfo.IsDir() {
			files, err := ioutil.ReadDir(filename)
			if err != nil {
				return fmt.Errorf("目录 %s 读属文件时错：%v", filename, err)
			}
			for _, file := range files {
				if path.Ext(file.Name()) == ".json" {
					if err := save(path.Join(filename, file.Name())); err != nil {
						return err
					}
				}
			}
		} else if path.Ext(filename) == ".json" {
			if err := save(filename); err != nil {
				return err
			}
		}
	}
	return nil
}
