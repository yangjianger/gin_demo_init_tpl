package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	//文件输出位置
	writeSyncer := getLogWriter()
	//文件格式
	encoder := getEncoder()

	//解析字符串级别
	level := zap.AtomicLevel{}
	if err := level.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		//默认info
		level = zap.NewAtomicLevel()
	}

	//core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//core := zapcore.NewCore(encoder, writeSyncer, level)

	//根据不同的模式，把日志输出到不同的位置
	var core zapcore.Core
	if viper.GetString("app.mode") == gin.DebugMode {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		//可以指定多个日志位置
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			//终端输出
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	//zap.AddCaller 显示行数
	logger := zap.New(core, zap.AddCaller())

	//替换全局loger
	zap.ReplaceGlobals(logger)

	//sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	//return zapcore.AddSync(file)
	//日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    //日志文件大小 单位 M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //备份时间 单位：天
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
