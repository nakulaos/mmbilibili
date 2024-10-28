package initializer

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
)

type LoggerInitializer struct {
	h             *server.Hertz
	logLevel      string
	logFileName   string
	logMaxSize    int
	logMaxBackups int
	logMaxAge     int
}

func NewLoggerInitializer(h *server.Hertz, logFileName string, logLevel string, logMaxSize int, logMaxBackups int, logMaxAge int) *LoggerInitializer {
	return &LoggerInitializer{h: h, logFileName: logFileName, logLevel: logLevel, logMaxSize: logMaxSize, logMaxBackups: logMaxBackups, logMaxAge: logMaxAge}
}

func (l *LoggerInitializer) Init() error {
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(logLevel(l.logLevel))

	//asyncWriter := &zapcore.BufferedWriteSyncer{
	//	WS: zapcore.AddSync(&lumberjack.Logger{
	//		Filename:   l.logFileName,
	//		MaxSize:    l.logMaxSize,
	//		MaxBackups: l.logMaxBackups,
	//		MaxAge:     l.logMaxAge,
	//	}),
	//	FlushInterval: time.Minute,
	//}
	//hlog.SetOutput(asyncWriter)

	//// 在应用关闭时同步日志
	//l.h.OnShutdown = append(l.h.OnShutdown, func(ctx context.Context) {
	//	asyncWriter.Sync()
	//})

	return nil
}

func MustInitLogger(h *server.Hertz, logFileName string, logLevel string, logMaxSize int, logMaxBackups int, logMaxAge int) {
	l := NewLoggerInitializer(h, logFileName, logLevel, logMaxSize, logMaxBackups, logMaxAge)
	if err := l.Init(); err != nil {
		panic(err)
	}
}

func logLevel(level string) hlog.Level {
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
