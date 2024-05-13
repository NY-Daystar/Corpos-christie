package main

import (
	"log"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/core"
	"github.com/NY-Daystar/corpos-christie/user"
)

// Configuration of the application
var cfg *config.Config

// Init configuration file
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // get line and file log
	cfg = config.New()
}

// Launching program
func main() {
	var user *user.User = new(user.User)
	core.Start(cfg, user)
}
