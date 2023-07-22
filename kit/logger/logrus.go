package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,                  // 日志颜色
		DisableTimestamp: false,                 // 是否禁用时间戳
		FullTimestamp:    true,                  // 是否显示完整时间
		TimestampFormat:  "2006-01-02 15:04:05", // 时间格式
		QuoteEmptyFields: true,                  // 空字段括在引号中
	})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.InfoLevel)
}
