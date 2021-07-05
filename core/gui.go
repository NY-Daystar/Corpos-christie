package core

import (
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/user"
)

// Program in GUI
type GUIMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// Start Core program in GUI Mode
func (m GUIMode) start() bool {
	return false
}
