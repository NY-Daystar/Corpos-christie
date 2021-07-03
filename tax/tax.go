package tax

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
	"github.com/LucasNoga/corpos-christie/user"

	"github.com/olekukonko/tablewriter"
)

// Result from processing income
type Result struct {
	income      int          // Input income from the user
	tax         float64      // Tax to pay from the user
	remainder   float64      // Value Remain for the user
	taxTranches []TaxTranche // List of tax by tranches
}

// Struct to catch tax capture for each tranche
type TaxTranche struct {
	tax     float64        // Tax in € on a tranche for the user
	tranche config.Tranche // Param of this tranche (Min, Max, Rate)
}

// Start tax calculator
// Calculate from income seized by user
func StartTaxCalulator(cfg *config.Config, user *user.User) {
	var status bool = true
	// Ask income's user
	fmt.Print("1. Enter your income (Revenu net imposable): ")
	_, err := user.AskIncome()
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
	result := calculateTax(user, cfg)

	// Show user
	user.Show()

	// Ask user if he wants to see tax tranches
	if ok, err := user.AskTaxDetails(); ok {
		if err != nil {
			log.Printf("Error: asking tax details, details: %v", err)
		}
		showTaxTranche(result)
	}

	if status {
		fmt.Println(colors.Green("Tax process successful"))
	} else {
		fmt.Println(colors.Red("Tax process failed"))
	}
	fmt.Println("--------------------------------------------------------------")

	// ask user to restart program else we exit
	fmt.Print("Would you want to enter a new income (Y/n): ")
	if user.AskRestart() {
		fmt.Println("Restarting program...")
		StartTaxCalulator(cfg, user)
	} else {
		fmt.Println("Quitting tax_calculator")
	}
}

// Start reverse tax calculator
// Calculate income needed from tax estimated after seized remainder income
//TODO a faire
func StartRevTaxCalulator(cfg *config.Config, user *user.User) {

	// TODO Boss final (v0.0.9)
	// TODO Pour l'exercice final on prendra le problème en sens inverse et on permettra à
	// TODO l'utilisateur d'entrer la somme désiré (après impôt) et le système calculera les
	// TODO revenus net à avoir pour obtenir cette somme après l'imposition.

	var status bool = true
	// Ask income's user
	fmt.Print("1. Enter your income wished after taxes (Revenu après impot): ")
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
	result := calculateReverseTax(user, cfg)

	// Show user
	user.Show()

	// Ask user if he wants to see tax tranches
	if ok, err := user.AskTaxDetails(); ok {
		if err != nil {
			log.Printf("Error: asking tax details, details: %v", err)
		}
		showTaxTranche(result)
	}

	if status {
		fmt.Println(colors.Green("Tax process successful"))
	} else {
		fmt.Println(colors.Red("Tax process failed"))
	}
	fmt.Println("--------------------------------------------------------------")

	// ask user to restart program else we exit
	fmt.Print("Would you want to enter a new income (Y/n): ")
	if user.AskRestart() {
		fmt.Println("Restarting program...")
		StartTaxCalulator(cfg, user)
	} else {
		fmt.Println("Quitting tax_calculator")
	}
}

// Processing the tax to pay from the income
func calculateTax(user *user.User, cfg *config.Config) Result {
	var tax float64
	var imposable float64 = float64(user.Income)
	user.CalculateParts()

	// if user has parts then its imposable is divided by parts number
	if user.Parts != 0 {
		imposable /= user.Parts
	}

	// Store each tranche taxes
	var taxTranches []TaxTranche = make([]TaxTranche, 0)

	// for each tranche
	for _, tranche := range cfg.Tax.Tranches {
		var taxTranche = calculateTranche(imposable, tranche)
		taxTranches = append(taxTranches, taxTranche)

		// add into final tax the tax tranche
		tax += taxTranche.tax
	}

	// if user has parts then its tax are multiplied by parts number after we calculate tax
	if user.Parts != 0 {
		tax *= user.Parts
	}

	// Format to round in integer tax and remainder
	result := Result{
		income:      user.Income,
		tax:         math.Round(tax),
		remainder:   float64(user.Income) - math.Round(tax),
		taxTranches: taxTranches,
	}

	// Add data into the user
	user.Tax = result.tax
	user.Remainder = result.remainder

	return result
}

// Processing the tax to pay from the remainder that the user want to get at the end
func calculateReverseTax(user *user.User, cfg *config.Config) Result {

	//TODO calcualte Tax
	// Income on arrive a calculer les tax
	// Remainder = Tax - Income

	//TODO calculate Rev Tax
	// Remainder on arrive a calculer les tax
	// Income = Remainder + Tax

	var tax float64
	var remainder float64 = float64(user.Remainder)
	// log.Printf("tax %v, remainder %v", tax, remainder)
	user.CalculateParts()

	// log.Printf("User %+v", user)

	// // if user has parts then its imposable is divided by parts number
	// if user.Parts != 0 {
	// 	imposable /= user.Parts
	// }

	// Store each tranche taxes
	var taxTranches []TaxTranche = make([]TaxTranche, 0)

	// // for each tranche
	for _, tranche := range cfg.Tax.Tranches {
		var taxTranche = calculateReverseTranche(remainder, tranche)
		log.Printf("Tax tranche %v", taxTranche)
		taxTranches = append(taxTranches, taxTranche)

		// add into final tax the tax tranche
		tax += taxTranche.tax
	}

	// // if user has parts then its tax are multiplied by parts number after we calculate tax
	// if user.Parts != 0 {
	// 	tax *= user.Parts
	// }

	// Format to round in integer tax and remainder
	result := Result{
		income:      int(remainder) + int(tax),
		tax:         math.Round(tax),
		remainder:   user.Remainder,
		taxTranches: taxTranches,
	}

	// Add data into the user
	user.Tax = result.tax
	user.Income = result.income

	return result
}

// Calculate tax for each tranche of your imposable
func calculateTranche(imposable float64, tranche config.Tranche) TaxTranche {
	var taxTranche TaxTranche = TaxTranche{
		tranche: tranche,
	}

	// convert rate string like '10%' into float 10.00
	rate, _ := utils.ConvertPercentageToFloat64(tranche.Rate)

	// if income is superior to maximum of the tranche to pass to tranch superior
	if int(imposable) > tranche.Max {
		taxTranche.tax = float64(tranche.Max-tranche.Min) * (rate / 100) // Diff between min and max of the tranche applied tax rate
	} else if int(imposable) > tranche.Min && int(imposable) < tranche.Max { // if your income is between min and max tranch is the last operation
		taxTranche.tax = float64(int(imposable)-tranche.Min) * (rate / 100) // Diff between min of the tranche and the income of the user,applied tax rate
	}
	return taxTranche
}

// Calculate reverse tax for each tranche from your remainder
func calculateReverseTranche(remainder float64, tranche config.Tranche) TaxTranche {
	var taxTranche TaxTranche = TaxTranche{
		tranche: tranche,
	}

	// convert rate string like '10%' into float 10.00
	rate, _ := utils.ConvertPercentageToFloat64(tranche.Rate)

	// if income is superior to maximum of the tranche to pass to tranch superior
	if int(remainder) > tranche.Max {
		taxTranche.tax = float64(tranche.Max-tranche.Min) * (rate / 100) // Diff between min and max of the tranche applied tax rate}
	} else if int(remainder) > tranche.Min && int(remainder) < tranche.Max { // if your income is between min and max tranch is the last operation
		taxTranche.tax = float64(int(remainder)-tranche.Min) * (rate / 100) // Diff between min of the tranche and the income of the user,applied tax rate
	}

	return taxTranche
}

// Show every tax at each tranch
func showTaxTranche(result Result, args ...interface{}) {
	var highlighted bool = false // if you want highlight data in table

	// Test args
	if len(args) > 0 {
		if args[0].(string) != "" {
			highlighted = true
		}
	}

	// Crete data to append on the table
	var data [][]string
	for i, val := range result.taxTranches {
		var index int = i + 1

		var trancheNumber string = fmt.Sprintf("Tranche %d", index)
		var min string = fmt.Sprintf("%s €", strconv.Itoa(val.tranche.Min))
		var max string = fmt.Sprintf("%s €", strconv.Itoa(val.tranche.Max))
		rate, _ := utils.ConvertPercentageToFloat64(val.tranche.Rate)
		var rateStr string = fmt.Sprintf("%s %%", strconv.Itoa(int(rate)))
		var tax string = fmt.Sprintf("%s €", strconv.Itoa(int(val.tax)))

		var line []string = make([]string, 5)
		line[0] = trancheNumber
		line[1] = min
		line[2] = max
		line[3] = rateStr
		line[4] = tax
		data = append(data, line)
	}

	// Install this: $ go get https://github.com/olekukonko/tablewriter
	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(true) // Set Border to false

	// Setting header
	var header []string = []string{"Tranche", "Min", "Max", "Rate", "Tax"}
	table.SetHeader(header)

	// Add data and Highlights Data
	if highlighted {
		for _, row := range data {
			tax, _ := strconv.ParseInt(strings.TrimSpace(strings.TrimSuffix(row[4], "€")), 10, 64)

			// if tax > 0 € red color
			if tax > 0 {
				table.Rich(row, []tablewriter.Colors{
					{},
					{},
					{},
					{},
					{tablewriter.Bold, tablewriter.FgRedColor}})
			} else {
				table.Append(row)
			}
		}
	} else { //Add classy data
		table.AppendBulk(data)
	}

	// Add footer
	var footer []string = []string{
		"Result",
		"Remainder",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.remainder))),
		"Total Tax",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.tax))),
	}
	table.SetFooter(footer)

	fmt.Println("\t\t\t Tax Details \t\t\t")
	table.Render()
}
