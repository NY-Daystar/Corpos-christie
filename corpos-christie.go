package main

import (
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/core"
	"github.com/NY-Daystar/corpos-christie/gui/model"
)

// Entrypoint of the program
func main() {
	var cfg *config.Config = config.New()
	var user *model.User = new(model.User)
	core.Start(cfg, user)
}
