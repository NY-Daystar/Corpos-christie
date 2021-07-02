// Handle program command
package core

import (
	"fmt"
	"log"
	"os"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
)

// List of the commands in console mod
type Command struct {
	name        string                           // Name of the command
	command     string                           // Code name of the command
	description string                           // Description of the command
	exec        func(*config.Config, *user.User) // Function to execute command
}

// List of options in console mode
var OPTIONS = []Command{
	{
		name:        "tax_calculator",
		command:     "tax_calculator",
		exec:        func(cfg *config.Config, user *user.User) { tax.Start(cfg, user) },
		description: "Calculate your tax from your income (income > tax)",
	},
	{
		name:    "reverse_tax_calculator",
		command: "reverse_tax_calculator",
		exec: func(cfg *config.Config, user *user.User) {
			log.Println(colors.Yellow("Not implemented yet, comming soon"))
		},
		description: "[WIP] Estimate your income from a tax amount (tax > income)",
	},
	{
		name:    "tax_history",
		command: "tax_history",
		exec: func(cfg *config.Config, user *user.User) {
			log.Println(colors.Yellow("Not implemented yet, comming soon"))
		},
		description: "[WIP] Show history of tax calculator",
	},
	{
		name:    "db",
		command: "db",
		exec: func(cfg *config.Config, user *user.User) {
			log.Println(colors.Yellow("Not implemented yet, comming soon"))
		},
		description: "[WIP] Get Db information",
	},
	{
		name:        "options",
		command:     "options",
		exec:        func(cfg *config.Config, user *user.User) {},
		description: "Show options list",
	},
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

	// var modeGUI *ConsoleMode = new(ConsoleMode)

	var modeSelected string = selectMode(os.Args)

	switch m := modeSelected; m {
	case "GUI":
		mode = GUIMode{config: cfg, user: user}
	case "console":
		// mode = ConsoleMode{config: cfg, user: user}
	default:
		mode = GUIMode{config: cfg, user: user}

	}
	// launch program
	ok := mode.start()

	// if doesn't work launch console
	if !ok {
		log.Printf(colors.Red("mode '%s' ")+colors.Red("cannot be launched. Launch console mode"), colors.Yellow(modeSelected))
		// var modeConsole *ConsoleMode = new(ConsoleMode)

		// modeConsole.start()

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

// Start Core program in console mode
// Show options to user
// And redirect to good choice
func (mode ConsoleMode) start() bool {
	log.Printf("Project: %s", colors.Yellow(mode.config.Name))
	log.Printf("Version: %s", colors.Yellow(mode.config.Version))

	// Loop so start program until user wants to exit
	for {
		showOptions()
		//TODO trouver un package qui fait de l'autocompletion pour mes commnades 'tax_calculator' etc...
		var optionEntered string = mode.user.ChooseOption()

		testOption, cmd := verifyOption(optionEntered)

		// if option is invalid
		if !testOption {
			fmt.Printf(colors.Red("Invalid option '%s', try again\n"), optionEntered)
			continue
		}

		// Execute command if it's a command to execute
		mode.execCommand(cmd)
	}
}

// Launch function from Cmd
// return true if the command is not the end, false if we have to return to the main options
func (mode ConsoleMode) execCommand(cmd Command) {
	// if commande is to show options
	if cmd.name == "options" {
		showOptions()
		return
	}

	// it's another command we execute it
	cmd.exec(mode.config, mode.user)
}

// Show list of options
func showOptions() {
	// prepend example command
	fmt.Println(colors.Yellow("\t\t\t List of options"))
	var exCommand Command = Command{name: "Exemple Command", description: "Description"}

	// Get all keys from console options list
	// to get max length of index for padding
	var cmdsName []string = getOptionsName(OPTIONS)

	cmdsName = append([]string{exCommand.name}, cmdsName...)

	// Show example command
	fmt.Printf(colors.Black("- [%s] %s %s\n"), exCommand.name, utils.SetPadding(cmdsName, exCommand.name), exCommand.description)
	// Show each options
	for _, cmd := range OPTIONS {
		fmt.Printf("- [%s] %s %s\n", colors.Magenta(cmd.name), utils.SetPadding(cmdsName, cmd.name), colors.Teal(cmd.description))
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
		}
	}
	return false, Command{}
}
