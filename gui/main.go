package gui

import (
	"errors"
	"log"
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/controller"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/view"
	"github.com/NY-Daystar/corpos-christie/user"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TODO voir si tout peut etre dans le package gui (model, view et controller)

// MVC for appliaction
type GUI struct {
	Model      *model.GUIModel
	View       *view.GUIView
	Controller *controller.GUIController
	Logger     *zap.Logger
}

// Start Launch GUI application
func Start(config *config.Config, user *user.User) {
	var logger = initLogger()
	logger.Info("Launch application")

	var model = model.NewModel(config, user, logger)
	var view = view.NewView(model, logger)
	var controller = controller.NewController(model, view, logger)

	var gui *GUI = &GUI{
		Model:      model,
		View:       view,
		Controller: controller,
	}

	gui.View.Start()
}

// initLogger create logger with zap librairy
func initLogger() *zap.Logger {
	configZap := zap.NewProductionEncoderConfig()
	configZap.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(configZap)

	// Create logs folder if not exists
	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

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
