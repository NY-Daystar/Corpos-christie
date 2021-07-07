// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"fmt"
	"os"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/user"
)

// Mode define the program mode (Console or GUI)
type Mode interface {
	start() bool // Start program in mode GUI or console
}

// Start Core program
// Get Options passed on program and launch appropriate system
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

// selectMode Check args passed in launch
// returns which mode to launch between GUI or console
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
