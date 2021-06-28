// Main package to start program
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
)

// Configuration of the application
var cfg *config.Config

// Start tax calculator from input user
func start(cfg *config.Config, user *user.User) bool {
	// Ask income's user
	_, err := user.AskIncome()
	if err != nil {
		log.Printf("Error: asking income for user, details: %v", err)
		return false
	}

	// Ask if user is in couple
	_, err = user.AskIsInCouple()
	if err != nil {
		log.Printf("Error: asking is in couple for user, details: %v", err)
		return false
	}

	// Ask if user hasChildren
	_, err = user.AskHasChildren()
	if err != nil {
		log.Printf("Error: asking has children, details: %v", err)
		return false
	}

	// Calculate tax
	result := tax.Process(user, cfg)

	// Show user
	user.Show()

	// Ask user if he wants to see tax tranches
	if ok, err := user.AskTaxDetails(); ok {
		if err != nil {
			log.Printf("Error: asking tax details, details: %v", err)
		}
		tax.ShowTaxTranche(result)
	}

	return true
}

// Init configuration file
func init() {
	cfg = new(config.Config)
	_, err := cfg.LoadConfiguration("./config.json")
	if err != nil {
		log.Printf(colors.Red("Unable to load config.json file, details: %v"), colors.Red(err))
		cfg.LoadDefaultConfiguration()
	}

	// get line and file log
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Starting program
func main() {
	log.Printf("Project: %v", colors.Yellow(cfg.Name))
	log.Printf("Version: %v", colors.Yellow(cfg.Version))

	// Init user
	var user *user.User = new(user.User)

	// Loop so start program until user wants to exit
	for {
		status := start(cfg, user)
		if status {
			log.Println(colors.Green("Tax process successful"))
		} else {
			log.Println(colors.Red("Tax process failed"))
		}
		fmt.Println("--------------------------------------------------------------")

		// ask user to restart program else we exit
		if user.AskRestart() {
			log.Printf("Restarting program...")
			continue
		}
		break
	}

	log.Printf("Program exited...")
	os.Exit(0)
}
