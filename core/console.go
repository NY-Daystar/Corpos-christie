// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

package core

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
	"github.com/LucasNoga/corpos-christie/utils"
	"github.com/LucasNoga/corpos-christie/utils/colors"
)

// Console represents the console application
type Console struct {
	Config *config.Config // Config to use correctly the program
	User   *user.User     // User param to use program
}

// Command define a command to use in console app
type Command struct {
	index       int                                       // Number to type to exec command
	name        string                                    // Name of the command
	description string                                    // Description of the command
	exec        func(cfg *config.Config, user *user.User) // Function to execute command
}

// OPTIONS is the list of options usable in console app
var OPTIONS []Command

// Init OPTIONS variables
func init() {
	OPTIONS = []Command{
		{
			name:        "tax_calculator",
			exec:        tax.StartTaxCalculator,
			description: "Calculate your tax from your incomes (income > tax)",
		},
		{
			name:        "reverse_tax_calculator",
			exec:        tax.StartReverseTaxCalculator,
			description: "Estimate your incomes from a tax amount (tax > income)",
		},
		{
			name:        "show_tax_tranche",
			exec:        func(cfg *config.Config, user *user.User) { tax.ShowTaxTranche(*cfg) },
			description: "Show the scale of taxes from the year selected",
		},
		{
			name:        "show_tax_year_list",
			exec:        func(cfg *config.Config, user *user.User) { tax.ShowTaxList(*cfg) },
			description: "Show the list of years to calculate your taxes",
		},
		{
			name:        "show_tax_year_used",
			exec:        func(cfg *config.Config, user *user.User) { tax.ShowTaxListUsed(*cfg) },
			description: "Show the year base to calculate your taxes",
		},
		{
			name:        "select_tax_year",
			exec:        func(cfg *config.Config, user *user.User) { tax.SelectTaxYear(cfg) },
			description: "Select a tax year if you want to calculate your taxes based on metrics of another year",
		},
		{
			name:        "options",
			exec:        func(cfg *config.Config, user *user.User) { showOptions() },
			description: "Show options list",
		},
		{
			name:        "about",
			exec:        func(cfg *config.Config, user *user.User) { showAbout() },
			description: "Show options list",
		},
		{
			name:        "quit",
			exec:        func(cfg *config.Config, user *user.User) { fmt.Println("Quitting program"); os.Exit(0) },
			description: "Quit program",
		},
	}

	// Insert index command
	for i, v := range OPTIONS {
		if v.name == "quit" {
			continue
		}
		OPTIONS[i].index = i + 1
	}
}

// Start launch application in console
func (app Console) Start() {
	fmt.Printf("Project: %s\n", colors.Yellow(app.Config.Name))
	fmt.Printf("Version: %s\b", colors.Yellow(app.Config.Version))

	// Loop so start program until user wants to exit
	for {
		// Show options to user
		showOptions()

		var optionEntered string = chooseOption()

		optionVerified, cmd := verifyOption(optionEntered)

		// if option doesn't exists
		if !optionVerified {
			fmt.Printf(colors.Red("Invalid option")+"'%s'. "+colors.Red("Try again\n"), colors.Yellow(optionEntered))
			continue
		}

		// If option is valid we execute the associate command
		cmd.exec(app.Config, app.User)
		fmt.Println("----------------------------------------")

		time.Sleep(700 * time.Millisecond)

	}
}

// showOptions show in the console the list of options which can be selected
func showOptions() {
	// prepend example command
	fmt.Println(colors.Yellow("\t\t\t List of options"))
	var exCommand Command = Command{index: 0, name: "Exemple Command", description: "Description"}

	// Get all keys from console options list
	// to get max length of index for padding
	var cmdsName []string = getOptionsName(OPTIONS)

	cmdsName = append([]string{exCommand.name}, cmdsName...)

	// Show example command
	fmt.Printf(colors.Black("- [%d] - [%s] %s %s\n"), exCommand.index, exCommand.name, utils.SetPadding(cmdsName, exCommand.name), exCommand.description)
	// Show each options
	for _, cmd := range OPTIONS {
		fmt.Printf("- [%s] - [%s] %s %s\n", colors.Black(cmd.index), colors.Magenta(cmd.name), utils.SetPadding(cmdsName, cmd.name), colors.Teal(cmd.description))
	}
	fmt.Println()
}

// showAbout show in the console the description of the application
func showAbout() {
	fmt.Printf("Application name: %s\n", colors.Yellow(config.APP_NAME))
	fmt.Println("Description: Application to calculate taxes in France developped in Golang")
	fmt.Printf("GitHub: %s\n", colors.Yellow(config.APP_LINK))
	fmt.Printf("Version : %s\n", colors.Yellow(fmt.Sprintf("v%s", config.APP_VERSION)))
	fmt.Printf("Author: %s\n", colors.Yellow(config.APP_AUTHOR))

}

// getOptionsName returns in a list the name of the commands in OPTIONS variable
func getOptionsName(cmds []Command) []string {
	var list []string = make([]string, 0, len(OPTIONS))
	for _, cmd := range cmds {
		list = append(list, cmd.name)
	}
	return list
}

// verifyOption check if option entered by a user is valid
// return true if the option exists, false if not
// if option exist return the Command struct associate to the option name
func verifyOption(optionEntered string) (bool, Command) {
	for _, cmd := range OPTIONS {
		// if the option choose is the name or the index of the command
		if cmd.name == optionEntered || strconv.Itoa(cmd.index) == optionEntered {
			return true, cmd
		}
	}
	return false, Command{}
}

// chooseOption ask to the user which command he wants to execute in console
// returns string seized in console by the user (define the command name)
func chooseOption() string {
	fmt.Print(colors.Green("Type an option > "))
	var input string = utils.ReadValue()
	return input
}
