package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// Settings data store in settings file
type Settings struct {
	logger   *zap.Logger
	Theme    int     `json:"theme"`
	Language *string `json:"language"`
	Currency *string `json:"currency"`
	Year     *string `json:"year"`
	Smtp     *Smtp   `json:"smtp"`
}

// outgoing mail server to send mail
type Smtp struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Load gui settings from settings file
func Load(logger *zap.Logger, filePath string) (Settings, error) {
	var settings Settings
	settings.logger = logger

	if filePath == "" {
		filePath = utils.GetSettingsFile()
	}

	settingsPath, _ := filepath.Abs(filePath)
	settings.logger.Info("Loading settings", zap.String("path", settingsPath))

	settingsFile, err := os.Open(settingsPath)
	if err != nil {
		settings.logger.Warn("Settings file error: ", zap.String("error", err.Error()))
		settings.logger.Info("Create and load default settings")
		return createDefaultSettings(), nil
	}

	jsonParser := json.NewDecoder(settingsFile)

	if err := jsonParser.Decode(&settings); err != nil {
		settings.logger.Fatal("Can't decode json : ", zap.String("error", err.Error()))
	}

	if err := settings.isValid(); err != nil {
		settings.logger.Error("Settings invalid : ", zap.String("error", err.Error()))
		return createDefaultSettings(), nil
	}

	return settings, settingsFile.Close()
}

// createDefaultSettings create settings file with default value
func createDefaultSettings() Settings {
	var settingsDefault = Settings{
		Theme:    GetDefaultTheme(),
		Language: GetDefaultLanguage(),
		Currency: GetDefaultCurrency(),
		Year:     GetDefaultYear(),
		Smtp:     GetDefaultSmtp(),
	}
	file, _ := json.MarshalIndent(settingsDefault, "", " ")
	_ = os.WriteFile(utils.GetSettingsFile(), file, 0644)
	return settingsDefault
}

// isConform verify is settings parse from file is valid
func (s *Settings) isValid() error {
	var errors = make([]string, 0)

	if s.Language == nil {
		errors = append(errors, "La langue est introuvable dans la configuration")
	}
	if s.Currency == nil {
		errors = append(errors, "La devise monnétaire est introuvable dans la configuration")
	}
	if s.Smtp == nil {
		errors = append(errors, "La connexion SMTP de la configuration est introuvable")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
}

// Set change value of data and write file with settings data
func (s *Settings) Set(key string, value interface{}) {
	switch key {
	case "theme":
		s.Theme = value.(int)
	case "language":
		v := value.(string)
		s.Language = &v
	case "currency":
		v := value.(string)
		s.Currency = &v
	case "year":
		v := value.(string)
		s.Year = &v
	}
	s.save()
}

// Save write file with settings data
func (s *Settings) save() {
	settingsPath, err := filepath.Abs(utils.GetSettingsFile())
	if err != nil {
		s.logger.Error("Can't get absolute path of settings", zap.String("error", err.Error()))
	}
	file, _ := json.MarshalIndent(s, "", " ")
	err = os.WriteFile(settingsPath, file, 0644)
	if err != nil {
		s.logger.Error("Save settings", zap.String("error", err.Error()))
	}
}
