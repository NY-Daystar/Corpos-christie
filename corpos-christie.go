package main

import (
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/core"
	"github.com/NY-Daystar/corpos-christie/user"
)

// Entrypoint of the program
func main() {
	var cfg *config.Config = config.New()
	var user *user.User = new(user.User)
	core.Start(cfg, user)
}
