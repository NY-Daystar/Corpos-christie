package main

import (
	"context"
	"corpos-christie/go/config"
	"corpos-christie/go/model"
	"corpos-christie/go/settings"
	"corpos-christie/go/tax"
	"corpos-christie/go/utils"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// App struct
type App struct {
	ctx    context.Context
	Config *config.Config
}

func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	var logger = initLogger()

	logger.Debug("Start Updater")
	//path, err := updater.StartUpdater() 	// TODO a reactiver
	// logger.Sugar().Debugf("Chemin: %v\n", path)
	// logger.Sugar().Errorf("Error: %v\n", err)
	logger.Debug("End Updater")

	a.Config = config.New()
	a.ctx = ctx
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

// fonction exposée à React
func (a *App) Calculate(income string) float64 {
	incString, _ := utils.ConvertStringToFloat64(income)
	var inc int = int(incString)
	user := &model.User{Income: inc} // TODO a mettre dans la methode
	cfg := config.New()              // TODO a mettre dans la methode
	result := tax.CalculateTax(user, cfg)

	return result.Tax
}

// TODO a check
func (a *App) GetHistory() []model.History {
	return []model.History{
		{Date: "2025-01-01", Income: 45000, Couple: false, IsInCouple: "no", Children: 0},
		{Date: "2025-02-01", Income: 50000, Couple: true, IsInCouple: "yes", Children: 2},
	}
}

func (a *App) GetCurrencies() []string {
	return settings.GetCurrencies()
}

func (a *App) GetDefaultCurrency() *string {
	return settings.GetDefaultCurrency()
}

func (a *App) GetLanguage() []string {
	return settings.GetCurrencies()
}

func (a *App) GetYears() []string {
	return settings.GetYears(a.Config)
}

func (a *App) GetDefaultYear() *string {
	return settings.GetDefaultYear()
}
