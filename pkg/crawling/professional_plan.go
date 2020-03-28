package crawling

// ProfessionalPlan represents information returned from the crawler abou professional plan
type ProfessionalPlan struct {
	Charge         string `json:"charge"`
	TransferTax    string `json:"transfer_tax"`
	MonthlyPayment string `json:"monthly_payment"`
}
