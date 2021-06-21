package main

import (
	"corpos-christie/colors"
	"corpos-christie/utils"
	"fmt"
	"log"
	"os"
)

// Start tax calculator from input user
func start() bool {
	fmt.Print("Enter your income (Revenu net imposable): ")
	var input string = utils.ReadValue()
	r, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax income is not convertible in int, details: %v", err)
		return false
	}

	log.Printf("income to calculate tax: %v", colors.Red(r))

	//TODO faire un fichier process pour faire les calculs d'impot
	//TODO Creer un fichier de config.json avec une struct tranches et une struct tranche qui gere ca
	//TODO Dans cette struct on a un min en int, un max en int et un pourcentage en int
	//TODO Tranche de revenu jusqu'à 10 084 € imposée à 0 % = 0 €
	//TODO Tranche de revenu de 10 085 € à 25 710 € : soit 15 625 € imposée à 11 % : 15 625 € x 11 % = 1 718,75 €
	//TODO Tranche de revenu de 25 711 € à 73 516 € imposée à 30 % : soit 6 289 € (obtenu en effectuant le calcul 32 000 -  25 711) x 30 % = 1 886,7 €.

	//TODO Calculer l'imposition
	//TODO Faire ensuite le differentiel pour savoir ce qui nous reste

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

func main() {
	log.Printf("Project: %v", colors.Yellow("Corpos-Christie")) //TODO gerer le nom de projet dans un fichier de config
	log.Printf("Version %v", colors.Yellow("1.0.0"))            //TODO gerer la version dynamiquemnt

	var keep bool
	for ok := true; ok; ok = keep {
		status := start()
		log.Printf("Status of operation: %v", status)
		keep = askRestart()
	}

	log.Printf("Program exited...")
	os.Exit(0)
}
