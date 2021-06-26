package user

import (
	"corpos-christie/colors"
	"corpos-christie/utils"
	"errors"
	"fmt"
	"log"
)

// Ask income to the user if ok return true, else return false
func (user *User) AskIncome() (bool, error) {
	fmt.Print("1. Enter your income (Revenu net imposable): ")
	var input string = utils.ReadValue()
	income, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax income is not convertible in int, details: %v", err)
		return false, err
	}
	user.Income = income
	return true, nil
}

// Ask if the user is in couple, ok return true, else return false
func (user *User) AskIsInCouple() (bool, error) {
	fmt.Print("2. Are you in couple (Y/n) ? ")
	var input string = utils.ReadValue()
	if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
		user.IsInCouple = true
	} else if input == "" || input == "N" || input == "n" || input == "No" || input == "no" {
		user.IsInCouple = false
	} else {
		return false, errors.New("invalid response you have to answer by (yes/Yes/Y/y or no/No/N/n)")
	}
	return true, nil
}

// Ask if the user does have children and how many
func (user *User) AskHasChildren() (bool, error) {
	fmt.Print("3.2 How many children do you have ? ")
	var input string = utils.ReadValue()

	// user can skip the question
	if input == "" {
		return true, nil
	}

	childrens, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Childrens value is not convertible in int, details: %v", err)
		return false, err
	}

	user.Children = childrens
	return true, nil
}

// Ask if the user does have children and how many
func (user *User) AskTaxDetails() (bool, error) {
	fmt.Print("Do you want to see tax details (Y/n) ? ")
	var input string = utils.ReadValue()
	if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
		return true, nil
	} else if input == "" || input == "N" || input == "n" || input == "No" || input == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid response you have to answer by (yes/Yes/Y/y or no/No/N/n)")
	}
}

// Processing to calculate part of the user
func (user *User) CalculateParts() {
	var parts float64
	// if user is in couple we have 2 parts
	if user.IsInCouple {
		parts = 2
	}

	// for each child of the user we put 0.5 parts
	parts += float64(user.Children) * 0.5
	user.Parts = parts
}

// Show user fields
func (user *User) Show() {
	var isInCouple = "No"
	if user.IsInCouple {
		isInCouple = "Yes"
	}
	fmt.Printf("Income:\t\t%v €\n", colors.Red(user.Income))
	fmt.Printf("In couple:\t%v\n", colors.Red(isInCouple))
	fmt.Printf("Children:\t%v\n", colors.Red(user.Children))
	fmt.Printf("Parts:\t\t%v\n", colors.Red(user.Parts))
	fmt.Printf("Tax:\t\t%v €\n", colors.Green(user.Tax))
	fmt.Printf("Remainder:\t%v €\n", colors.Green(user.Remainder))
}
