package logger

import (
	"RiYu/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// LogData 日志数据结构
type LogData struct {
	Time    time.Time         `bson:"时间"`
	Level   string            `bson:"等级"`
	Content string            `bson:"报文"`
	Data    map[string]string `bson:"数据,omitempty"`
}

// 日志等级
const (
	Info  = "Info"
	Fatal = "Fatal"
)

// Logger 日志机
type Logger struct {
	collection *mongo.Collection
}

// New 新日志服务
func New(client *mongodb.Client, dbName, colName string) *Logger {
	return &Logger{client.Database(dbName).Collection(colName)}
}

// Log 记日志
func (logger *Logger) Log(log *LogData) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Time = time.Now()
	_, err = logger.collection.InsertOne(ctx, log)
	return
}
