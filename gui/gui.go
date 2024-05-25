package gui

import (
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/user"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// MVC for appliaction
type GUI struct {
	Model      *GUIModel
	View       *GUIView
	Controller *GUIController
	Logger     *zap.Logger
}

// Start Launch GUI application
// TODO a commenter
func Start(config *config.Config, user *user.User, display ...bool) {
	var logger = initLogger()
	logger.Info("Launch application")

	var model = NewModel(config, user, logger)
	var view = NewView(model, logger)
	var controller = NewController(model, view, logger)

	var gui *GUI = &GUI{
		Model:      model,
		View:       view,
		Controller: controller,
	}

	// Launch GUI if bool ok
	if len(display) == 0 {
		gui.View.Start(controller)
	}
}

// initLogger create logger with zap librairy
func initLogger() *zap.Logger {
	configZap := zap.NewProductionEncoderConfig()
	configZap.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(configZap)

	// Create logs folder if not exists
	path := "logs"
	os.Mkdir(path, os.ModePerm)

	logger := lumberjack.Logger{
		Filename:   config.LOGS_PATH, // File path
		MaxSize:    500,              // 500 megabytes per files
		MaxBackups: 3,                // 3 files before rotate
		MaxAge:     15,               // 15 days
	}

	writer := zapcore.AddSync(&logger)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, zapcore.DebugLevel),
	)

	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	log.Info("Zap logger set",
		zap.String("path", logger.Filename),
		zap.Int("filesize", logger.MaxSize), zap.Int("backupfile", logger.MaxBackups),
		zap.Int("fileage", logger.MaxAge),
	)

	return log
}
