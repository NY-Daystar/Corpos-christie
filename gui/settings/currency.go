package settings

// Enum for currency
const (
	EURO   string = "€"
	DOLLAR string = "$"
	LIVRE  string = "£"
)

// GetCurrencies get array of all currencies
func GetCurrencies() []string {
	return []string{EURO, DOLLAR, LIVRE}
}
