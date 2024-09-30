package core

import (
	"os"
	"path"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/updater"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Start Core program
// Get Options passed on program and launch appropriate system
func Start(cfg *config.Config, user *model.User) {
	var logger = initLogger()

	logger.Debug("Start Updater")
	path, err := updater.StartUpdater()
	logger.Sugar().Debugf("Chemin: %v\n", path)
	logger.Sugar().Errorf("Error: %v\n", err)
	logger.Debug("End Updater")

	gui.Start(cfg, user, logger, path)
}

// initLogger create logger with zap librairy
func initLogger() *zap.Logger {
	configZap := zap.NewProductionEncoderConfig()
	configZap.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(configZap)

	// Create logs folder if not exists
	appPath, _ := utils.GetAppDataPath()
	var logsFolder = path.Join(appPath, config.APP_NAME, "logs")
	os.Mkdir(logsFolder, os.ModePerm)

	logger := lumberjack.Logger{
		Filename:   utils.GetLogsFile(), // File path
		MaxSize:    500,                 // 500 megabytes per files
		MaxBackups: 3,                   // 3 files before rotate
		MaxAge:     15,                  // 15 days
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
