package settings

const (
	DARK  string = "Dark"
	LIGHT string = "Light"
)

// ThemeYml Yaml struct for theme's app
type ThemeYaml struct {
	Dark  string `yaml:"dark"`
	Light string `yaml:"light"`
}

// GetTheme Get value of last theme selected
func GetDefaultTheme() string {
	// TODO get value from .setting file
	// TODO log debug to show change theme
	return DARK
}
