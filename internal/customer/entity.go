package customer

type CustomerHistory struct {
	CPF                      string
	IsValidCPF               bool
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

func NewCustomerHistory(
	CPF string,
	IsValidCPF bool,
	Private, Incomplete *bool,
	LastPurchaseDate *string,
	AverageTicket, LastPurchaseTicket *float64,
	MostFrequentStore, LastPurchaseStore *string,
	IsValidMostFrequentStore, IsValidLastPurchaseStore bool,
) CustomerHistory {
	return CustomerHistory{
		CPF:                      CPF,
		IsValidCPF:               IsValidCPF,
		Private:                  Private,
		Incomplete:               Incomplete,
		LastPurchaseDate:         LastPurchaseDate,
		AverageTicket:            AverageTicket,
		LastPurchaseTicket:       LastPurchaseTicket,
		MostFrequentStore:        MostFrequentStore,
		IsValidMostFrequentStore: IsValidMostFrequentStore,
		LastPurchaseStore:        LastPurchaseStore,
		IsValidLastPurchaseStore: IsValidLastPurchaseStore,
	}
}
