package settings

// Handle the themes in GUI settings

import (
	"reflect"
)

// ThemeYml Yaml struct for theme's app
type ThemeYaml struct {
	Dark  string `yaml:"dark"`
	Light string `yaml:"light"`
}

// GetTheme Get value of last theme selected
func GetTheme() string {
	// TODO get value from .setting file
	// TODO log debug to show change theme
	return "Dark"
}

// GetThemes Parse ThemeYaml struct to get value of each field
func (t ThemeYaml) GetThemes() []string {
	v := reflect.ValueOf(t)
	themes := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		themes[i] = v.Field(i).String()
	}
	return themes
}
