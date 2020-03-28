package crawling

// ProfessionalPlan represents informations about professional plan
type ProfessionalPlan struct {
	Charge         string `json:"charge"`
	TransferTax    string `json:"transfer_tax"`
	MonthlyPayment string `json:"monthly_payment"`
}
