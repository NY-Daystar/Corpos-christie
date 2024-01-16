// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core defines core functions to run GUI app or Console app
package core

import (
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui"
	"github.com/NY-Daystar/corpos-christie/user"
)

// Enum for launched mode
const (
	GUI     string = "gui"
	CONSOLE string = "console"
)

// Start Core program
// Get Options passed on program and launch appropriate system
func Start(cfg *config.Config, user *user.User) {
	var appSelected string = selectMode(os.Args)

	// Launch program (Console or GUI)
	switch m := appSelected; m {
	case GUI:
		gui.GUI{Config: cfg, User: user}.Start()
	case CONSOLE:
		Console{Config: cfg, User: user}.Start()
	default:
		gui.GUI{Config: cfg, User: user}.Start()
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
