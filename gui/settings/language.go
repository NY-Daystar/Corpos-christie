package settings

// Handle the languages in GUI settings

import (
	"reflect"
)

// Languages Yaml struct for theme's app
type LanguageYaml struct {
	English string `yaml:"english"`
	French  string `yaml:"french"`
}

// LanguageYaml Yaml struct to get language data
type Yaml struct {
	Code         string       // code of the language (fr, en, etc...)
	Theme        ThemeYaml    `yaml:"themes"`
	Languages    LanguageYaml `yaml:"languages"`
	File         string       `yaml:"file"`
	Settings     string       `yaml:"settings"`
	Income       string       `yaml:"income"`
	Status       string       `yaml:"status"`
	Children     string       `yaml:"children"`
	Tax          string       `yaml:"tax"`
	Remainder    string       `yaml:"remainder"`
	Share        string       `yaml:"share"`
	SaveTax      string       `yaml:"save_tax"`
	ThemeCode    string       `yaml:"theme"`
	LanguageCode string       `yaml:"language"`
	Help         string       `yaml:"help"`
	About        string       `yaml:"about"`
	Author       string       `yaml:"author"`
	Close        string       `yaml:"close"`
	Quit         string       `yaml:"quit"`
}

// GetLanguage Get value of last language selected (fr, en)
func GetLanguage() string {
	// TODO get value from .setting file
	// TODO log debug to show change language
	return "en"
}

// GetLanguages Parse LanguagesXml struct to get value of each field
func (l LanguageYaml) GetLanguages() []string {
	v := reflect.ValueOf(l)
	languages := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		languages[i] = v.Field(i).String()
	}
	return languages
}
