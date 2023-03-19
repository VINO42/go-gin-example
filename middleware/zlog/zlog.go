package zlog

import (
	"github.com/natefinch/lumberjack"
	"github.com/vino42/go-gin-example/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// set time format
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// set log codec -> json
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	file, _ := os.OpenFile(setting.LogHome, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	fileWriterSyncer := zapcore.AddSync(file)

	core := zapcore.NewTee(
		// output to console
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		// output to file
		zapcore.NewCore(encoder, fileWriterSyncer, zapcore.DebugLevel),
	)

	logger = zap.New(core)
}

// log rotate
func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	// use lumberjack to rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "D:\\demo1\\src\\demo\\demo06\\zap-log\\logs\\zap.log",
		MaxSize:    10, // filesize, MB
		MaxBackups: 7,  // >7, than rotate
		MaxAge:     1,  // frequency, 1 day
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {

	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
