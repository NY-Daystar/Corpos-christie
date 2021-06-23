package main

import (
	"corpos-christie/colors"
	"corpos-christie/config"
	"corpos-christie/core"
	"corpos-christie/utils"
	"fmt"
	"log"
	"os"
)

var cfg *config.Config

// Start tax calculator from input user
func start(cfg *config.Config) bool {
	fmt.Print("Enter your income (Revenu net imposable): ")
	var input string = utils.ReadValue()
	income, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax income is not convertible in int, details: %v", err)
		return false
	}

	// calculate tax
	result := core.Process(income, cfg)
	fmt.Printf("Income:\t\t%v €\nTax:\t\t%v €\nRemainder:\t%v €\n", colors.Red(result.Income), colors.Red(result.Tax), colors.Red(result.Remainder))

	return true
}

// Ask user if he wants to restart program
func askRestart() bool {
	for {
		fmt.Print("Would you want to enter a new income (Y/n): ")
		var input string = utils.ReadValue()
		if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
			log.Printf("Restarting program...")
			return true
		} else {
			return false
		}
	}
}

// Init configuration file
func init() {
	cfg = new(config.Config)
	cfg.LoadConfiguration("./config.json")

	// get line and file log
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Starting program
func main() {
	log.Printf("Project: %v", colors.Yellow(cfg.Name))
	log.Printf("Version: %v", colors.Yellow(cfg.Version))

	var keep bool
	for ok := true; ok; ok = keep {
		status := start(cfg)
		log.Printf("Status of operation: %v", status)
		fmt.Println("--------------------------------------------------------------")
		keep = askRestart()
	}

	log.Printf("Program exited...")
	os.Exit(0)
}
