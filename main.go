// Main package to start program
package main

import (
	"log"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/core"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/user"
)

const (
	APP_NAME    string = "Corpos-Christie"
	APP_VERSION string = "0.0.8"
)

// Configuration of the application
var cfg *config.Config

// Init configuration file
func init() {
	cfg = new(config.Config)
	cfg.Name = APP_NAME
	cfg.Version = APP_VERSION
	_, err := cfg.LoadConfiguration("./config.json")
	if err != nil {
		log.Printf(colors.Red("Unable to load config.json file, details: %s"), colors.Red(err))
		cfg.LoadDefaultConfiguration()
	}

	// get line and file log
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Starting program
func main() {
	// Init user
	var user *user.User = new(user.User)

	core.Start(cfg, user)
}
