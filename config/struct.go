package config

// Define the program config
type Config struct {
	Name    string
	Version string
	Tax     Tax
	TaxList []Tax
}

// Define the tranche on a specific year
type Tax struct {
	Year     int
	Tranches []Tranche
}

// Define one of the tranch of tax
type Tranche struct {
	Min        int
	Max        int
	Percentage float64
}
