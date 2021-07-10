// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package uses store function to interact with user
package user

import (
	"errors"
	"fmt"
	"log"

	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
)

// User defines a the user of the program
type User struct {
	Income     int     // Income (Revenu imposable) of the user
	Tax        float64 // Tax to pay for the user
	Remainder  float64 // Money remind after tax paid
	Shares     float64 // Shares (or Parts in french) is the family quotient base on if you are in couple and if you have children to adjust your taxes
	IsInCouple bool    // User is he in couple or not
	Children   int     // number of children of the user
}

// AskIncome asks the income of the user to calculate tax and set it into user struct
// if value is set returns true, otherwise false
func (user *User) AskIncome() (bool, error) {
	var input string = utils.ReadValue()
	income, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax income is not convertible in int, details: %v", err)
		return false, err
	}
	user.Income = income
	return true, nil
}

// AskRemainder asks the remainder of the user to calculate reverse tax and set it into user struct
// if value is set returns true, otherwise false
func (user *User) AskRemainder() (bool, error) {
	var input string = utils.ReadValue()
	remainder, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax remainder is not convertible in int, details: %v", err)
		return false, err
	}
	user.Remainder = float64(remainder)
	return true, nil
}

// AskIsInCouple asks if the user is in couple set it into user struct
// returns response of the user
func (user *User) AskIsInCouple() (bool, error) {
	response, err := askYesNo()
	if err != nil {
		return false, err
	}
	if response {
		user.IsInCouple = true
	} else {
		user.IsInCouple = false
	}
	return response, nil
}

// AskHasChildren asks the number of children of the user and set it into user struct
// if value is set returns true, otherwise false
func (user *User) AskHasChildren() (bool, error) {
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

// AskTaxDetails asks to the user if he wants to see details of his taxes
// returns true if wants otherwise false
func (user *User) AskTaxDetails() (bool, error) {
	fmt.Print("Do you want to see tax details (Y/n) ? ")
	response, err := askYesNo()
	if err != nil {
		return false, err
	}
	return response, nil
}

// AskRestart asks the user if he wants to retry a calculation of tax
// returns true if wants otherwise false
func (user *User) AskRestart() bool {
	response, _ := askYesNo()
	return response
}

// GetShares returns the shares of the user
func (user *User) GetShares() float64 {
	return user.Shares
}

// Show show details of the user struct
func (user *User) Show() {
	var isInCouple = "No"
	if user.IsInCouple {
		isInCouple = "Yes"
	}
	fmt.Printf("Income:\t\t%s €\n", colors.Red(user.Income))
	fmt.Printf("In couple:\t%s\n", colors.Red(isInCouple))
	fmt.Printf("Children:\t%s\n", colors.Red(user.Children))
	fmt.Printf("Shares:\t\t%s\n", colors.Red(user.Shares))
	fmt.Printf("Tax:\t\t%s €\n", colors.Green(user.Tax))
	fmt.Printf("Remainder:\t%s €\n", colors.Green(user.Remainder))
}

// askYesNo handle the interaction of the user if he has to answer by 'yes' or 'no'
// returns true if the user say 'yes', false if he answered 'no'
// returns an error if the seize is not interpretable
func askYesNo() (bool, error) {
	var input string = utils.ReadValue()
	if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
		return true, nil
	} else if input == "" || input == "N" || input == "n" || input == "No" || input == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid response you have to answer by (yes/Yes/Y/y or no/No/N/n)")
	}
}
