package customer

type CustomerHistory struct {
	CPF                      string
	IsCPFValid               bool
	Private                  *bool
	Incomplete               *bool
	LastPurchaseDate         *string
	AverageTicket            *float64
	LastPurchaseTicket       *float64
	MostFrequentStore        *string
	IsValidMostFrequentStore bool
	LastPurchaseStore        *string
	IsValidLastPurchaseStore bool
}
