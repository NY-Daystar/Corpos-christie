// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package tax is the algorithm to calculate taxes
package tax

import (
	"math"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
)

// Result define the result after calculating tax
type Result struct {
	Income      int          // Input income from the user
	Tax         float64      // Tax to pay from the user
	Remainder   float64      // Value Remain for the user
	TaxTranches []TaxTranche // List of tax by tranches
	Shares      float64      // family quotient to adjust taxes (parts in french)
}

// TaxTranche represent the tax calculating for each tranch when we calculate tax
type TaxTranche struct {
	Tax     float64        // Tax in € on a tranche for the user
	Tranche config.Tranche // Param of the tranche calculated (Min, Max, Rate)
}

// calculateTax determine the tax to pay from the income of the user
// returns the result of the processing
func CalculateTax(user *model.User, cfg *config.Config) Result {
	var tax float64
	var taxable = float64(user.Income)
	var shares = getShares(*user)

	// Divide taxable by shares
	taxable /= shares

	// Store each tranche taxes
	var taxTranches = make([]TaxTranche, 0)

	// for each tranche
	for _, tranche := range cfg.GetTax().Tranches {
		var taxTranche = calculateTranche(int(taxable), tranche)
		taxTranches = append(taxTranches, taxTranche)

		// add into final tax the tax tranche
		tax += taxTranche.Tax
	}

	// Reajust tax by shares
	tax *= shares

	// Format to round in integer tax and remainder
	result := Result{
		Income:      user.Income,
		Tax:         math.Round(tax),
		Remainder:   float64(user.Income) - math.Round(tax),
		TaxTranches: taxTranches,
		Shares:      shares,
	}

	// Add data into the user
	user.Tax = result.Tax
	user.Remainder = result.Remainder
	user.Shares = result.Shares

	return result
}

// calculateReverseTax determine the income to have, and tax to pay from the remainder of the user
// returns the result of the processing
func CalculateReverseTax(user *model.User, cfg *config.Config) Result {
	var income float64

	var taxTranches []TaxTranche
	var shares = getShares(*user)

	var incomeAfterTaxes = user.Remainder
	var target = incomeAfterTaxes // income to find

	// Divide taxable by shares
	target /= shares

	// Brut force to find target with incomeAfterTaxes
	for {

		var tax float64

		taxTranches = make([]TaxTranche, 0)
		// for each tranche
		for _, tranche := range cfg.GetTax().Tranches {
			var taxTranche = calculateTranche(int(target), tranche)
			taxTranches = append(taxTranches, taxTranche)

			// add into final tax the tax tranche
			tax += taxTranche.Tax
		}

		tax *= shares

		// When target has been reached
		if incomeAfterTaxes <= target*shares-tax {
			income = target*shares - shares
			break
		}
		// Increase target to find if we not find
		target++
	}

	// Format to round in integer tax and remainder
	result := Result{
		Income:      int(income),
		Tax:         math.Round(income - incomeAfterTaxes),
		Remainder:   incomeAfterTaxes,
		TaxTranches: taxTranches,
		Shares:      shares,
	}

	// Add data into the user
	user.Income = result.Income
	user.Tax = result.Tax

	return result
}

// calculateTranche calculate the tax for the tranche base on your taxable income
// returns TaxTranche which amount to pay for the specific tranche
func calculateTranche(taxable int, tranche config.Tranche) TaxTranche {
	var taxTranche = TaxTranche{
		Tranche: tranche,
	}

	// convert rate in percentage
	// Ex:'30' in 0.30 to get 30%
	var rate = float64(tranche.Rate) / 100.

	// If income is superior to maximum of the tranche to pass to tranch superior
	// Diff between min and max of the tranche applied tax rate
	if taxable > tranche.Max {
		taxTranche.Tax = float64(tranche.Max-tranche.Min) * rate
		// else if your income taxable is between min and max tranch is the last operation
		// Diff between min of the tranche and the income of the user applied tax rate
	} else if taxable > tranche.Min && taxable < tranche.Max {
		taxTranche.Tax = float64(int(taxable)-tranche.Min) * rate
	}
	return taxTranche
}

// getShares calculate the family quotient of the user (parts in french)
// returns the shares calculated
func getShares(user model.User) float64 {
	var shares float64 = 1 // single person only 1 share

	// if user is in couple we have 1 more shares,
	if user.IsInCouple {
		shares += 1
	}

	// if parent is single and have children it's a isolated parent
	if user.IsIsolated() {
		shares += 0.5
	}

	// For the two first children we add 0.5
	for i := 1; i <= user.Children && i <= 2; i++ {
		shares += 0.5
	}

	// For the others children we add 1
	for i := 3; i <= user.Children; i++ {
		shares += 1
	}

	// for each child of the user we put 0.5 shares
	return shares
}
