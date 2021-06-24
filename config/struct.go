package config

// Define the program config
type Config struct {
	Name     string    `json:"name"`
	Version  string    `json:"version"`
	Tranches []Tranche `json:"tranches"`
}

// Define one of the tranch of tax
type Tranche struct {
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Percentage float64 `json:"percentage"`
}
