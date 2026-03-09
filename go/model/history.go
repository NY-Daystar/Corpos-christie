package model

// Get History of tax
type History struct {
	Date       string `json:"date"`     // saving date
	Income     int    `json:"income"`   // Income seized by the user
	Couple     bool   `json:"couple"`   // Couple or not for this maths calculation
	IsInCouple string `json:"-"`        // Format bool `Couple with yes or no`
	Children   int    `json:"children"` // Number of children
}
