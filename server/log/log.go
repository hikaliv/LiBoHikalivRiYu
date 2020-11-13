package log

import (
	"RiYu/logger"
	"RiYu/mongodb"
)

/**
 * 全局单例
 */
var client *mongodb.Client
var log *logger.Logger

// Start 日志开机
func Start() (err error) {
	client, err = mongodb.Start()
	if err != nil {
		return
	}
	log = logger.New(client, "Japanese", "Log")
	return
}

// Log 日记
func Log(data *logger.LogData) error {
	return log.Log(data)
}

// Stop 日志停机
func Stop() error {
	return client.Stop()
}
