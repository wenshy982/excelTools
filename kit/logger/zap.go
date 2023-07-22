package logger

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() {
	// 创建一个配置为生产环境的 zap logger
	encoder := zapcore.NewConsoleEncoder(newEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	// 替换全局的 zap logger
	zap.ReplaceGlobals(logger)

	// 创建一个新的 logger，这个 logger 的输出已经被重定向到上面创建的 logger 里
	zapLog := zap.NewStdLog(logger)

	// 让标准库的 log 使用我们新创建的 zapLog
	log.SetFlags(0)
	log.SetOutput(zapLog.Writer())
}

// newEncoderConfig 获取zapcore.EncoderConfig
func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",                           // 自定义时间key
		LevelKey:      "level",                          // 自定义等级key
		NameKey:       "logger",                         // 自定义logger key
		CallerKey:     "caller",                         // 自定义调用者key
		MessageKey:    "msg",                            // 自定义消息key
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 报错等级大写加颜色
		StacktraceKey: "stacktrace",                     // 调用堆栈信息
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(time.DateTime))
		}, // 时间格式化
		EncodeDuration:   zapcore.SecondsDurationEncoder, // 时间格式化
		EncodeCaller:     zapcore.ShortCallerEncoder,     // FullCallerEncoder 显示完整文件路径
		ConsoleSeparator: " ",                            // 控制台输出分隔符
	}
}
