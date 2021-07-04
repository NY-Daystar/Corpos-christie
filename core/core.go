// Handle program command
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

// List of the commands in console mod
type Command struct {
	index       int                              // number to type to exec command
	name        string                           // Name of the command
	command     string                           // Code name of the command
	description string                           // Description of the command
	exec        func(*config.Config, *user.User) // Function to execute command
}

// List of options in console mode
var OPTIONS []Command

// Init Options variables
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
			name:    "tax_history",
			command: "tax_history",
			exec: func(cfg *config.Config, user *user.User) {
				fmt.Println(colors.Yellow("Not implemented yet, comming soon"))
			},
			description: "[WIP] Show history of tax calculator",
		},
		{
			name:    "db",
			command: "db",
			exec: func(cfg *config.Config, user *user.User) {
				fmt.Println(colors.Yellow("Not implemented yet, comming soon"))
			},
			description: "[WIP] Get Db information",
		},
		{
			name:    "history",
			command: "history",
			exec: func(cfg *config.Config, user *user.User) {
				fmt.Println(colors.Yellow("Not implemented yet, comming soon"))
			},
			description: "Show command history",
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

// Mode of program launc (Console or GUI)
type Mode interface {
	start() bool // Start program in mode GUI or console
}

// Program in GUI
type GUIMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// Program in console mode
type ConsoleMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// Start Core program
// Get Options passed on program and launch systems
func Start(cfg *config.Config, user *user.User) {
	var mode Mode

	var modeSelected string = selectMode(os.Args)

	switch m := modeSelected; m {
	case "GUI":
		mode = GUIMode{config: cfg, user: user}
	case "console":
		mode = ConsoleMode{config: cfg, user: user}
	default:
		mode = GUIMode{config: cfg, user: user}

	}
	// launch program
	ok := mode.start()

	// if doesn't work launch console mode
	if !ok {
		fmt.Printf(colors.Red("mode '%s' ")+colors.Red("cannot be launched. Launch console mode\n"), colors.Yellow(modeSelected))
		ConsoleMode{config: cfg, user: user}.start()
	}
}

// Return which mode to launch between GUI or console
func selectMode(args []string) string {
	// if no args specified launch GUI
	if len(args) < 2 {
		return "GUI"
	} else {
		var mode string = args[1]
		switch m := mode; m {
		case "--gui":
			return "GUI"
		case "--console":
			return "console"
		default:
			return "default"
		}
	}
}

// Start Core program in GUI Mode
func (m GUIMode) start() bool {
	return false
}

// Handle command from user to know how do you do in console mode
func chooseOption() string {
	fmt.Print(colors.Green("Type an option > "))
	var input string = utils.ReadValue()
	return input
}

// Start Core program in console mode
// Show options to user
// And redirect to good choice
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

// Launch function of the command
func (mode ConsoleMode) execCommand(cmd Command) {
	cmd.exec(mode.config, mode.user)
	fmt.Println("----------------------------------------")
}

// Show list of options
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

// Get only the name of the options for console mode
func getOptionsName(cmds []Command) []string {
	var list []string = make([]string, 0, len(OPTIONS))
	for _, cmd := range cmds {
		list = append(list, cmd.name)
	}
	return list
}

// Verify if option entered by a user is valid
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
