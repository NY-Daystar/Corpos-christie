package core

import (
	"corpos-christie/config"
	"math"
)

// Processing the tax to pay from the income
func Process(income int, cfg *config.Config) Result {
	var tax float64

	// for each tranche
	for _, tranche := range cfg.Tranches {
		// log.Printf("%v", tranche)
		if income > tranche.Max { // if income is superior to maximum of the tranche to pass to tranch superior
			// we get the diff between min and max of the tranch we apply the tax percentage and we add it to the user
			tax = float64(tranche.Max-tranche.Min) * (tranche.Percentage / 100)
			continue
		} else if income > tranche.Min && income < tranche.Max { // if your income is between min and max tranch is the last operation
			// we get the diff between min of the tranch and the income of the user, we applied percentage then we add it to the user
			tax += float64(income-tranche.Min) * (tranche.Percentage / 100)
			break
		}
	}

	// Format to round in integer tax and remainder
	result := Result{Income: income, Tax: math.Round(tax), Remainder: float64(income) - math.Round(tax)}

	return result
}
