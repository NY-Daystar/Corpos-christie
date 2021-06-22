package core

// Result from processing income
type Result struct {
	Income    int     //Input income from the user
	Tax       float64 // Tax to pay from the user
	Remainder float64 // Value Remain for the user
}
