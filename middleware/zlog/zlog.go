package zlog

import (
	"github.com/natefinch/lumberjack"
	"github.com/vino42/go-gin-example/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.SugaredLogger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// set time format
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// set log codec -> json
	//encoder := zapcore.NewJSONEncoder(encoderConfig)
	// set to nomarl log encod
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	file, _ := os.OpenFile(setting.LogHome+"/"+setting.LogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	fileWriterSyncer := zapcore.AddSync(file)

	core := zapcore.NewTee(
		// output to console
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		// output to file
		zapcore.NewCore(encoder, fileWriterSyncer, zapcore.DebugLevel),
	)

	logger = zap.New(core, zap.AddCaller()).Sugar()
}

// log rotate
func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	// use lumberjack to rotate
	//Filename: 日志文件的位置
	//MaxSize：在进行切割之前，日志文件的最大大小（以 MB 为单位）
	//MaxBackups：保留旧文件的最大个数
	//MaxAges：保留旧文件的最大天数
	//Compress：是否压缩 / 归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   setting.LogHome + "/" + setting.LogName,
		MaxSize:    10, // filesize, MB
		MaxBackups: 7,  // >7, than rotate
		MaxAge:     1,  // frequency, 1 day
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(message string, fields ...interface{}) {

	logger.Infof(" | "+message, fields...)
}

func Debug(message string, fields ...interface{}) {
	logger.Debugf(" | "+message, fields...)
}

func Error(message string, fields ...interface{}) {
	logger.Errorf(" | "+message, fields...)
}

func Warn(message string, fields ...interface{}) {
	logger.Warnf(" | "+message, fields...)
}
