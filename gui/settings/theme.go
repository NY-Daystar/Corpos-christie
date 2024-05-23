package settings

const (
	LIGHT int = 0
	DARK  int = 1
)

// ThemeYaml Yaml struct for theme's app
type ThemeYaml struct {
	Light string `yaml:"light"`
	Dark  string `yaml:"dark"`
}

// GetDefaultTheme Get value of last theme selected
func GetDefaultTheme() int {
	return LIGHT
}
