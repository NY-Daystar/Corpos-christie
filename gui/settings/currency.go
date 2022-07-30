package settings

// Handle the themes in GUI settings

// Enum for currency
const (
	EURO   string = "€"
	DOLLAR string = "$"
	LIVRE  string = "£"
)

// GetCurrency Get value of last currency
func GetCurrency() string {
	// TODO get value from .setting file
	// TODO log debug to show change theme
	return EURO
}

// GetCurrencies get array of all currencies
func GetCurrencies() []string {
	return []string{EURO, DOLLAR, LIVRE}
}
