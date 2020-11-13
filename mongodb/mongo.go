package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client 客端
type Client struct {
	*mongo.Client
}

// Start 服务启动
func Start() (cli *Client, err error) {
	// 定五秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 启联客户端，试五秒
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	// 五秒之内，试联通与否
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}
	cli = &Client{client}
	return
}

// Stop 服务中止
func (client *Client) Stop() error {
	return client.Disconnect(context.Background())
}
