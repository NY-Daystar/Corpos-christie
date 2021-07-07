// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/user"
)

// GUIMode represents the program parameters to launch in console mode the application
type GUIMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// start launch core program in GUI Mode
func (m GUIMode) start() bool {
	return false
}
