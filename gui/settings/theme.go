package settings

const (
	DARK  int = 0
	LIGHT int = 1
)

// ThemeYaml Yaml struct for theme's app
type ThemeYaml struct {
	Dark  string `yaml:"dark"`
	Light string `yaml:"light"`
}

// GetTheme Get value of last theme selected
func GetDefaultTheme() int {
	return DARK
}
