// Tax package
package tax

import (
	"testing"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/user"
)

// For testing
// $ cd tax
// $ go test -v

// Global variables
var CONFIG *config.Config

// Init global variables
func init() {
	CONFIG = new(config.Config)
	CONFIG.Tax = config.Tax{
		Year: 2021,
		Tranches: []config.Tranche{
			{Min: 0, Max: 10084, Rate: "0%"},
			{Min: 10085, Max: 25710, Rate: "11%"},
			{Min: 25711, Max: 73516, Rate: "30%"},
			{Min: 73517, Max: 158122, Rate: "41%"},
			{Min: 158123, Max: 1000000, Rate: "45%"},
		},
	}
}

// Calculate tax for a single person with 32000 of income
func TestCalculateTaxForSinglePerson(t *testing.T) {
	user := user.User{
		Income: 32000,
	}

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 32000, tax: 3605, remainder: 28395}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate tax for a couple with 2 children, testing parts with a couple and 2 childrens
func TestCalculateTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Income:     55950,
		IsInCouple: true,
		Children:   2,
		Parts:      3,
	}

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 55950, tax: 2826, remainder: 53124}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate reverse tax for a single person to get at the end 28395
func aTestCalculateReverseTaxForSinglePerson(t *testing.T) {
	user := user.User{
		Remainder: 28395,
	}

	result := calculateReverseTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 32000, tax: 3605, remainder: 28395}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate reverse tax for a couple with 2 children, testing parts with a couple and 2 childrens
func aTestCalculateReverseTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Remainder:  53124,
		IsInCouple: true,
		Children:   2,
		Parts:      3,
	}

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 55950, tax: 2826, remainder: 53124}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}
