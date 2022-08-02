package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/LucasNoga/corpos-christie/config"
	"go.uber.org/zap"
)

// Settings data store in settings file
type Settings struct {
	logger   *zap.Logger
	Theme    int    `json:"theme"`
	Language string `json:"language"`
	Currency string `json:"currency"`
}

// Load gui settings from settings file
func Load(logger *zap.Logger) Settings {
	var settings Settings
	settings.logger = logger
	settingsPath, _ := filepath.Abs(config.SETTINGS_PATH)
	settings.logger.Info("Loading settings", zap.String("path", settingsPath))

	settingsFile, err := os.Open(settingsPath)
	if err != nil {
		settings.logger.Warn("Settings file error: ", zap.String("error", err.Error()))
		settings.logger.Info("Create and load default settings")
		return createDefaultSettings()
	}
	defer settingsFile.Close()
	jsonParser := json.NewDecoder(settingsFile)
	if err := jsonParser.Decode(&settings); err != nil {
		settings.logger.Fatal("Can't decode json : ", zap.String("error", err.Error()))
	}
	return settings
}

// createDefaultSettings create settings file with default value
func createDefaultSettings() Settings {
	var settingsDefault Settings = Settings{
		Theme:    GetDefaultTheme(),
		Language: GetDefaultLanguage(),
		Currency: GetDefaultCurrency(),
	}
	file, _ := json.MarshalIndent(settingsDefault, "", " ")
	_ = ioutil.WriteFile(config.SETTINGS_PATH, file, 0644)
	return settingsDefault
}

// Set change value of data and write file with settings data
func (s *Settings) Set(key string, value interface{}) {
	switch key {
	case "theme":
		s.Theme = value.(int)
	case "language":
		s.Language = value.(string)
	case "currency":
		s.Currency = value.(string)
	}
	s.save()
}

// Save write file with settings data
func (s *Settings) save() {
	settingsPath, err := filepath.Abs(config.SETTINGS_PATH)
	if err != nil {
		s.logger.Error("Can't get absolute path of settings", zap.String("error", err.Error()))
	}
	file, _ := json.MarshalIndent(s, "", " ")
	err = ioutil.WriteFile(settingsPath, file, 0644)
	if err != nil {
		s.logger.Error("Save settings", zap.String("error", err.Error()))
	}
}
