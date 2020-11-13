package log

import "RiYu/logger"

// PanicError 异常处理
func PanicError(err interface{}) {
	var content string
	switch err.(type) {
	case string:
		content = err.(string)
	case error:
		content = err.(error).Error()
	default:
		content = "err 类型异常，报错丢失"
	}
	// 两级错误处理，应提供最后一级错误用以应对系统错误日志执行报错兜底
	if logerr := log.Log(&logger.LogData{
		Level:   logger.Fatal,
		Content: content,
	}); logerr != nil {
		fatal(logerr)
	}
}

func fatal(err error) {}
