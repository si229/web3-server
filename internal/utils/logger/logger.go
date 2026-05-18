package logger

import (
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log zerolog.Logger
var WatchLog zerolog.Logger

func Init() {
	logFile := &lumberjack.Logger{
		Filename:   "log/app.log", // 日志文件名
		MaxSize:    10,            // 每个文件最大 10MB
		MaxBackups: 5,             // 保留 5 个旧文件
		MaxAge:     7,             // 保留 7 天
		Compress:   true,          // 是否压缩 gzip
	}
	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
	Log = zerolog.New(multi).
		With().
		Timestamp().
		Caller().
		Logger()

	// 特殊日志（单独文件）
	WatchLogFile := &lumberjack.Logger{
		Filename:   "log/watchlog.log",
		MaxSize:    20,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}
	WatchLog = zerolog.New(WatchLogFile).With().Timestamp().Logger()

}
