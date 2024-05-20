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

// GetDefaultCurrency get value of default currency
func GetDefaultCurrency() *string {
	var currency string = EURO
	return &currency
}
