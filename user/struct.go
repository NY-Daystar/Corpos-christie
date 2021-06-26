package user

// Define a user
type User struct {
	Income     int     // Income (Revenu imposable) of the user
	Tax        float64 // Tax to pay for the user
	Remainder  float64 // Money remind after tax paid
	Parts      float64 // Parts of the user (calculate from isInCouple, childre)
	IsInCouple bool    // User is he in couple or not
	Children   int     // number of children of the user
}
