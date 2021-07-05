// Handle program command
package core

import (
	"fmt"
	"os"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/user"
)

// Mode of program launch (Console or GUI)
type Mode interface {
	start() bool // Start program in mode GUI or console
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
