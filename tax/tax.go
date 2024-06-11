// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package tax is the algorithm to calculate taxes
package tax

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/user"
	"github.com/NY-Daystar/corpos-christie/utils"
	"github.com/NY-Daystar/corpos-christie/utils/colors"

	"github.com/olekukonko/tablewriter"
)

// TODO remove console code en v3.0.0

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
	tranche config.Tranche // Param of the tranche calculated (Min, Max, Rate)
}

// StartTaxCalculator calculate taxes from income seized by user
func StartTaxCalculator(cfg *config.Config, user *user.User) {
	fmt.Printf("The calculator is based on %s\n", colors.Teal(cfg.GetTax().Year))
	status := true
	// Ask income's user
	fmt.Print("1. Enter your income\n    (en) Taxable income\n    (fr) Revenus net imposable\n> ")
	_, err := user.AskIncome()
	if err != nil {
		log.Printf("Error: asking income for user, details: %v", err)
		status = false
		return
	}

	// Ask if user is in couple
	fmt.Print("2. Are you in couple (Y/n) ? ")
	_, err = user.AskIsInCouple()
	if err != nil {
		log.Printf("Error: asking is in couple for user, details: %v", err)
		status = false
		return
	}

	// Ask if user hasChildren
	fmt.Print("3. How many children do you have ? ")
	_, err = user.AskHasChildren()
	if err != nil {
		log.Printf("Error: asking has children, details: %v", err)
		status = false
		return
	}

	// Calculate tax
	result := CalculateTax(user, cfg)
	user.Shares = result.Shares

	user.Show()

	// Ask user if he wants to see tax tranches
	fmt.Print("Do you want to see tax details (Y/n) ? ")
	if ok, err := user.AskTaxDetails(); ok {
		if err != nil {
			log.Printf("Error: asking tax details, details: %v", err)
		}
		showTaxTrancheResult(result, cfg.Tax.Year)
	}

	if status {
		fmt.Println(colors.Green("Tax process successful"))
	} else {
		fmt.Println(colors.Red("Tax process failed"))
	}
	fmt.Println("----------------------------------------")

	// ask user to restart program else we exit
	fmt.Print("Would you want to enter a new income (Y/n): ")
	if user.AskRestart() {
		fmt.Println("Restarting program...")
		StartTaxCalculator(cfg, user)
	} else {
		fmt.Println("Quitting tax_calculator")
	}
}

// StartReverseTaxCalculator calculate income needed from remainder seized by user
func StartReverseTaxCalculator(cfg *config.Config, user *user.User) {
	fmt.Printf("The calculator is based on %s\n", colors.Teal(cfg.GetTax().Year))
	status := true

	// Ask income's user
	fmt.Print("1. Enter your income wished\n    (en) Income after taxes income\n    (fr) Revenus après impot\n> ")
	_, err := user.AskRemainder()
	if err != nil {
		log.Printf("Error: asking income for user, details: %v", err)
		status = false
	}

	// Ask if user is in couple
	fmt.Print("2. Are you in couple (Y/n) ? ")
	_, err = user.AskIsInCouple()
	if err != nil {
		log.Printf("Error: asking is in couple for user, details: %v", err)
		status = false
	}

	// Ask if user hasChildren
	fmt.Print("3. How many children do you have ? ")
	_, err = user.AskHasChildren()
	if err != nil {
		log.Printf("Error: asking has children, details: %v", err)
		status = false
	}

	// Calculate tax
	result := CalculateReverseTax(user, cfg)
	user.Shares = result.Shares

	user.Show()

	// Ask user if he wants to see tax tranches
	fmt.Print("Do you want to see tax details (Y/n) ? ")
	if ok, err := user.AskTaxDetails(); ok {
		if err != nil {
			log.Printf("Error: asking tax details, details: %v", err)
		}
		showTaxTrancheResult(result, cfg.Tax.Year)
	}

	if status {
		fmt.Println(colors.Green("Tax process successful"))
	} else {
		fmt.Println(colors.Red("Tax process failed"))
	}
	fmt.Println("----------------------------------------")

	// ask user to restart program else we exit
	fmt.Print("Would you want to enter a new income (Y/n): ")
	if user.AskRestart() {
		fmt.Println("Restarting program...")
		StartReverseTaxCalculator(cfg, user)
	} else {
		fmt.Println("Quitting tax_calculator")
	}
}

// calculateTax determine the tax to pay from the income of the user
// returns the result of the processing
func CalculateTax(user *user.User, cfg *config.Config) Result {
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

	return result
}

// calculateReverseTax determine the income to have, and tax to pay from the remainder of the user
// returns the result of the processing
func CalculateReverseTax(user *user.User, cfg *config.Config) Result {
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
		tranche: tranche,
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
func getShares(user user.User) float64 {
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

// showTaxTranche show details of calculation showing every tax at each tranche
func showTaxTrancheResult(result Result, year int) {

	// Install this: $ go get https://github.com/olekukonko/tablewriter
	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(true) // Set Border to false

	// Setting header
	var header = []string{"Tranche", "Min", "Max", "Rate", "Tax"}
	table.SetHeader(header)

	// Create data to append on the table
	var data [][]string
	for i, val := range result.TaxTranches {
		index := i + 1

		var trancheNumber = fmt.Sprintf("Tranche %d", index)
		var min = fmt.Sprintf("%d €", val.tranche.Min)
		var max = fmt.Sprintf("%d €", val.tranche.Max)
		var rate = fmt.Sprintf("%d %%", val.tranche.Rate)
		var tax = fmt.Sprintf("%d €", int(val.Tax))

		var line = make([]string, 5)
		line[0] = trancheNumber
		line[1] = min
		line[2] = max
		line[3] = rate
		line[4] = tax
		data = append(data, line)
	}

	// Add data in table
	table.AppendBulk(data)

	// Add footer
	var footer = []string{
		"Result",
		"Remainder",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.Remainder))),
		"Total Tax",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.Tax))),
	}
	table.SetFooter(footer)

	fmt.Println(colors.Yellow("\t\t\t Tax Details \t\t\t"))
	fmt.Printf("For an income of %s € in %s\n", colors.Teal(result.Income), colors.Teal(year))
	table.Render()
}

// ShowTaxList show in the console the list of year metrics
func ShowTaxList(cfg config.Config) {
	fmt.Println(colors.Yellow("Tax list year"))
	fmt.Println("-------------")
	for _, v := range cfg.TaxList {
		var year = strconv.Itoa(v.Year)
		if cfg.GetTax().Year == v.Year {
			year = "* " + colors.Green(v.Year)
		}
		fmt.Printf("%s\n", year)
	}
}

// ShowTaxTranche show in the console the list of year metrics
func ShowTaxTranche(cfg config.Config) {
	fmt.Printf("Tax tranche of year (%s)\n", colors.Teal(cfg.GetTax().Year))
	fmt.Println("-------------")
	for index, tranche := range cfg.GetTax().Tranches {
		fmt.Printf("%d - From %s to %s - Rate taxes : %s\n", index, colors.Teal(tranche.Min), colors.Teal(tranche.Max), colors.Yellow(tranche.Rate))
	}
}

// ShowTaxListUsed show the current tax used in the console
func ShowTaxListUsed(cfg config.Config) {
	fmt.Printf("The tax year base to calculate your taxes is %s\n", colors.Teal(cfg.GetTax().Year))
}

// SelectTaxYear ask in console if you want
// Ask to the user if he wants to change the year of the tax metrics
// to calculate taxes from another year
func SelectTaxYear(cfg *config.Config) {
	fmt.Printf("The calculator is based on %s\n", colors.Teal(cfg.GetTax().Year))

	// Asking year
	fmt.Print("List of years: ")
	for _, v := range cfg.TaxList {
		var year = strconv.Itoa(v.Year)
		if cfg.GetTax().Year == v.Year {
			year = colors.Green(v.Year)
		}
		fmt.Printf("%s ", year)
	}
	fmt.Print("\nWhich year do you want ? ")

	var input = utils.ReadValue()

	year, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax year is not convertible in int, details: %v", err)
		return
	}

	cfg.ChangeTax(year)
	fmt.Printf("The tax year is now based on %s\n", colors.Teal(cfg.GetTax().Year))

}
