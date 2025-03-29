package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	//
	// sugar.Infof("Hello name: %s, age: %d", "ngductoann", 25) // like fmt.Printf(format, a)
	//
	// // Logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "ngductoann"), zap.Int("age", 25)) // type key-value
	//
	// // Development mode
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development mode")
	//
	// // Production mode
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production mode")

	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core)
	logger.Info("Info log", zap.Int("line", 1))
	logger.Info("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encodeConfig.TimeKey = "time"

	// From info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
