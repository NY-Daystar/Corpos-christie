// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
)

// ConsoleMode represents the program parameters to launch in console mode the application
type ConsoleMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// Command define a command to use in console mode
type Command struct {
	index       int                              // number to type to exec command
	name        string                           // Name of the command
	command     string                           // Code name of the command
	description string                           // Description of the command
	exec        func(*config.Config, *user.User) // Function to execute command
}

// OPTIONS is the list of options usable in console mode
var OPTIONS []Command

// Init OPTIONS variables
func init() {
	OPTIONS = []Command{
		{
			name:        "tax_calculator",
			command:     "tax_calculator",
			exec:        func(cfg *config.Config, user *user.User) { tax.StartTaxCalculator(cfg, user) },
			description: "Calculate your tax from your incomes (income > tax)",
		},
		{
			name:        "reverse_tax_calculator",
			command:     "reverse_tax_calculator",
			exec:        func(cfg *config.Config, user *user.User) { tax.StartReverseTaxCalculator(cfg, user) },
			description: "Estimate your incomes from a tax amount (tax > income)",
		},
		{
			name:        "show_tax_year_list",
			command:     "show_tax_year_list",
			exec:        func(cfg *config.Config, user *user.User) { tax.ShowTaxList(*cfg) },
			description: "Show the list of years to calculate your taxes",
		},
		{
			name:        "show_tax_year_used",
			command:     "show_tax_year_used",
			exec:        func(cfg *config.Config, user *user.User) { tax.ShowTaxListUsed(*cfg) },
			description: "Show the year base to calculate your taxes",
		},
		{
			name:        "select_tax_year",
			command:     "select_tax_year",
			exec:        func(cfg *config.Config, user *user.User) { tax.SelectTaxYear(cfg, user) },
			description: "Select a tax year if you want to calculate your taxes based on metrics of another year",
		},
		{
			name:        "options",
			command:     "options",
			exec:        func(cfg *config.Config, user *user.User) { showOptions() },
			description: "Show options list",
		},
		{
			name:        "quit",
			command:     "quit",
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

// start launch core program in console mode
// Show options to user in the console
// Interact with the user depending of the option selected
func (mode ConsoleMode) start() bool {
	fmt.Printf("Project: %s\n", colors.Yellow(mode.config.Name))
	fmt.Printf("Version: %s\b", colors.Yellow(mode.config.Version))

	// Loop so start program until user wants to exit
	for {
		showOptions()

		var optionEntered string = chooseOption()

		optionVerified, cmd := verifyOption(optionEntered)

		// if option doesn't exists
		if !optionVerified {
			fmt.Printf(colors.Red("Invalid option")+"'%s'. "+colors.Red("Try again\n"), colors.Yellow(optionEntered))
			continue
		}

		// If option is valid we execute the associate command
		mode.execCommand(cmd)

		time.Sleep(700 * time.Millisecond)

	}
}

// execCommand execute the function set in the exec field of the cmd
func (mode ConsoleMode) execCommand(cmd Command) {
	cmd.exec(mode.config, mode.user)
	fmt.Println("----------------------------------------")
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
		// if it's the command
		if cmd.command == optionEntered {
			return true, cmd
		} else if strconv.Itoa(cmd.index) == optionEntered { // if it's an index of command
			fmt.Printf("You selected the commmand %s\n", colors.Yellow(cmd.name))
			return true, cmd
		}
	}
	return false, Command{}
}

// chooseOption ask to the user which command he wants to execute in console mode
// returns string seized in console mode by the user (define the command name)
func chooseOption() string {
	fmt.Print(colors.Green("Type an option > "))
	var input string = utils.ReadValue()
	return input
}
