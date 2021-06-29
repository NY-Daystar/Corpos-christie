package tax

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/user"

	"github.com/olekukonko/tablewriter"
)

// Result from processing income
type Result struct {
	income      int          //Input income from the user
	tax         float64      // Tax to pay from the user
	remainder   float64      // Value Remain for the user
	taxTranches []TaxTranche // List of tax by tranches
}

// Struct to catch tax capture for each tranche
type TaxTranche struct {
	tax     float64        // Tax in € on a tranche for the user
	tranche config.Tranche // Param of this tranche (Min, Max, Percentage)
}

// Processing the tax to pay from the income
func Process(user *user.User, cfg *config.Config) Result {
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
		var taxTranche TaxTranche = TaxTranche{
			tranche: tranche,
		}

		// if income is superior to maximum of the tranche to pass to tranch superior
		if int(imposable) > tranche.Max {
			taxTranche.tax = float64(tranche.Max-tranche.Min) * (tranche.Percentage / 100) // Diff between min and max of the tranche applied tax percentage
		} else if int(imposable) > tranche.Min && int(imposable) < tranche.Max { // if your income is between min and max tranch is the last operation
			taxTranche.tax = float64(int(imposable)-tranche.Min) * (tranche.Percentage / 100) // Diff between min of the tranche and the income of the user,applied tax percentage
		}

		taxTranches = append(taxTranches, taxTranche)

		// add into final tax the tax tranche
		tax += taxTranche.tax
	}

	// if user has parts then its tax are multiplied by parts number
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

// Show every tax at each tranch
func ShowTaxTranche(result Result, args ...interface{}) {
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

		var line []string = make([]string, 5)
		line[0] = fmt.Sprintf("Tranche %v", index)
		line[1] = fmt.Sprintf("%v €", strconv.Itoa(val.tranche.Min))
		line[2] = fmt.Sprintf("%v €", strconv.Itoa(val.tranche.Max))
		line[3] = fmt.Sprintf("%v %%", strconv.Itoa(int(val.tranche.Percentage)))
		line[4] = fmt.Sprintf("%v €", strconv.Itoa(int(val.tax)))
		data = append(data, line)
	}

	// Install this: $ go get https://github.com/olekukonko/tablewriter
	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(true) // Set Border to false

	// Setting header
	var header []string = []string{"Tranche", "Min", "Max", "Percentage", "Tax"}
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
		fmt.Sprintf("%v €", strconv.Itoa(int(result.remainder))),
		"Total Tax",
		fmt.Sprintf("%v €", strconv.Itoa(int(result.tax))),
	}
	table.SetFooter(footer)

	fmt.Println("\t\t\t Tax Details \t\t\t")
	table.Render()
}
