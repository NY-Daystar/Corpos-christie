package core

import (
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui"
	"github.com/NY-Daystar/corpos-christie/user"
)

// Enum for launched mode
const (
	GUI       string = "gui"
	CONSOLE   string = "console"
	TEST_MODE string = "test"
)

// Start Core program
// Get Options passed on program and launch appropriate system
func Start(cfg *config.Config, user *user.User, mode ...string) {
	var appSelected string
	if len(mode) == 0 {
		appSelected = selectMode(os.Args)
	} else {
		appSelected = mode[0]
	}

	//fmt.Printf("appMode:  %+v\n", appSelected)
	// TODO envoyer le logger
	//updater.StartUpdater()
	//fmt.Printf("UPDATER TERMINE\n")

	// Launch program (Console or GUI)
	switch m := appSelected; m {
	case GUI:
		gui.Start(cfg, user)
	case CONSOLE:
		Console{Config: cfg, User: user}.Start()
	case TEST_MODE:
		return
	default:
		gui.Start(cfg, user)
	}
}

// selectMode Check args passed in launch
// returns which mode app to launch between GUI or console
func selectMode(args []string) string {
	// if no args specified launch GUI
	if len(args) < 2 {
		return GUI
	} else {
		var mode string = args[1]
		switch m := mode; m {
		case "--gui":
			return GUI
		case "--console":
			return CONSOLE
		default:
			return GUI
		}
	}
}
