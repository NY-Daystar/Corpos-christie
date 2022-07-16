// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package is the entrypoint of the program
package main

import (
	"log"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/core"
	"github.com/LucasNoga/corpos-christie/user"
)

// Configuration of the application
var cfg *config.Config

// Init configuration file
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // get line and file log

	// Setup config
	cfg = config.New()
}

// Launching program
func main() {
	// Init user
	var user *user.User = new(user.User)

	core.Start(cfg, user)
}
