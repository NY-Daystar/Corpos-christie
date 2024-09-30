// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package uses store function to interact with user
package model

// User defines a the user of the program
type User struct {
	Income     int     // Income (Revenu imposable) of the user
	Tax        float64 // Tax to pay for the user
	Remainder  float64 // Money remind after tax paid
	Shares     float64 // Shares (or Parts in french) is the family quotient base on if you are in couple and if you have children to adjust your taxes
	IsInCouple bool    // User is he in couple or not
	Children   int     // number of children of the user
}

// IsIsolated return bool if parent has children to raise alone
func (user *User) IsIsolated() bool {
	return !user.IsInCouple && user.Children > 0
}
