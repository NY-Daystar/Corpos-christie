package core

import (
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui"
	"github.com/NY-Daystar/corpos-christie/user"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Enum for launched mode
const (
	GUI       string = "gui"
	CONSOLE   string = "console"
	TEST_MODE string = "test"
)

// Start Core program
// Get Options passed on program and launch appropriate system
func Start(cfg *config.Config, user *user.User, mode ...string) {
	var logger = initLogger()
	var appSelected string
	if len(mode) == 0 {
		appSelected = selectMode(os.Args)
	} else {
		appSelected = mode[0]
	}

	//fmt.Printf("START UPDATER\n")
	//updater.StartUpdater(logger)
	//fmt.Printf("UPDATER TERMINE\n")

	// Launch program (Console or GUI)
	switch m := appSelected; m {
	case GUI:
		gui.Start(cfg, user, logger)
	case CONSOLE:
		Console{Config: cfg, User: user}.Start()
	case TEST_MODE:
		return
	default:
		gui.Start(cfg, user, logger)
	}
}

// selectMode Check args passed in launch
// returns which mode app to launch between GUI or console
func selectMode(args []string) string {
	// if no args specified launch GUI
	if len(args) < 2 {
		return GUI
	} else {
		var mode string = args[1]
		switch m := mode; m {
		case "--gui":
			return GUI
		case "--console":
			return CONSOLE
		default:
			return GUI
		}
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
