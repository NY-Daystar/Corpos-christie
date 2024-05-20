package settings

import (
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/utils"
)

// GetYears get array of all tax year
func GetYears(config *config.Config) []string {
	var years []string
	for _, tax := range config.TaxList {
		year := utils.ConvertIntToString(tax.Year)
		years = append(years, year)
	}
	return years
}

// GetDefaultYear get value of default year
func GetDefaultYear() *string {
	var config = config.New()
	years := GetYears(config)
	var year = years[len(years)-1]
	return &year
}
