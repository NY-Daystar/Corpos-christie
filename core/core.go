package core

import (
	"corpos-christie/config"
	"corpos-christie/user"
	"math"
)

// Processing the tax to pay from the income
func Process(user *user.User, cfg *config.Config) Result {
	var tax float64
	var imposable float64 = float64(user.Income)
	user.CalculateParts()

	// if user has parts then its imposable is divided by parts number
	if user.Parts != 0 {
		imposable /= user.Parts
	}

	// for each tranche
	for _, tranche := range cfg.Tranches {
		if int(imposable) > tranche.Max { // if income is superior to maximum of the tranche to pass to tranch superior
			// we get the diff between min and max of the tranch we apply the tax percentage and we add it to the user
			tax = float64(tranche.Max-tranche.Min) * (tranche.Percentage / 100)
			continue
		} else if int(imposable) > tranche.Min && int(imposable) < tranche.Max { // if your income is between min and max tranch is the last operation
			// we get the diff between min of the tranch and the income of the user, we applied percentage then we add it to the user
			tax += float64(int(imposable)-tranche.Min) * (tranche.Percentage / 100)
			break
		}
	}

	// if user has parts then its tax are multiplied by parts number
	if user.Parts != 0 {
		tax *= user.Parts
	}

	// Format to round in integer tax and remainder
	result := Result{Income: user.Income, Tax: math.Round(tax), Remainder: float64(user.Income) - math.Round(tax)}
	return result
}
