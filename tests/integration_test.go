package customer_tests

import (
	"bytes"
	"testing"

	c "github.com/lucasdanielgeo/customer-history-importer/internal/customer"
)

func TestImportationProcessIntegration(t *testing.T) {
	t.Run("Test importation process and check memory repository", func(t *testing.T) {
		repo := c.NewMemoryCustomerHistoryRepository()
		reader := bytes.NewBufferString(MockFileContent)
		customerImporter := c.CustomerHistoryImporter{
			Repository: repo,
			Reader:     reader,
		}
		customerImporter.Execute()

		for _, customerHistory := range MockCustomerHistorySlicePersisted {
			customerFromRepo, err := repo.Read(customerHistory.CPF)
			if err != nil {
				t.Errorf("error reading customer from repository: %v", err)
			}

			if customerFromRepo.CPF != customerHistory.CPF {
				t.Errorf("customer CPF field saved in memory repository does not match the mocked customer: %v != %v", customerFromRepo.CPF, customerHistory.CPF)
			}

			if customerFromRepo.IsValidCPF != customerHistory.IsValidCPF {
				t.Errorf("customer IsValidCPF field saved in memory repository does not match the mocked customer: %v != %v", customerFromRepo.IsValidCPF, customerHistory.IsValidCPF)
			}

			if customerFromRepo.IsValidMostFrequentStore != customerHistory.IsValidMostFrequentStore {
				t.Errorf("customer IsValidMostFrequentStore field saved in memory repository does not match the mocked customer: %v != %v", customerFromRepo.IsValidMostFrequentStore, customerHistory.IsValidMostFrequentStore)
			}

			if customerFromRepo.IsValidLastPurchaseStore != customerHistory.IsValidLastPurchaseStore {
				t.Errorf("customer IsValidLastPurchaseStore field saved in memory repository does not match the mocked customer: %v != %v", customerFromRepo.IsValidLastPurchaseStore, customerHistory.IsValidLastPurchaseStore)
			}

			if customerFromRepo.Private != nil && *customerFromRepo.Private != *customerHistory.Private {
				t.Errorf("customer Private field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.Private, *customerHistory.Private)
			}

			if customerFromRepo.Incomplete != nil && *customerFromRepo.Incomplete != *customerHistory.Incomplete {
				t.Errorf("customer Incomplete field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.Incomplete, *customerHistory.Incomplete)
			}

			if customerFromRepo.LastPurchaseDate != nil && *customerFromRepo.LastPurchaseDate != *customerHistory.LastPurchaseDate {
				t.Errorf("customer LastPurchaseDate field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.LastPurchaseDate, *customerHistory.LastPurchaseDate)
			}

			if customerFromRepo.AverageTicket != nil && *customerFromRepo.AverageTicket != *customerHistory.AverageTicket {
				t.Errorf("customer AverageTicket field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.AverageTicket, *customerHistory.AverageTicket)
			}

			if customerFromRepo.LastPurchaseTicket != nil && *customerFromRepo.LastPurchaseTicket != *customerHistory.LastPurchaseTicket {
				t.Errorf("customer LastPurchaseTicket field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.LastPurchaseTicket, *customerHistory.LastPurchaseTicket)
			}

			if customerFromRepo.MostFrequentStore != nil && *customerFromRepo.MostFrequentStore != *customerHistory.MostFrequentStore {
				t.Errorf("customer MostFrequentStore field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.MostFrequentStore, *customerHistory.MostFrequentStore)
			}

			if customerFromRepo.LastPurchaseStore != nil && *customerFromRepo.LastPurchaseStore != *customerHistory.LastPurchaseStore {
				t.Errorf("customer LastPurchaseStore field saved in memory repository does not match the mocked customer: %v != %v", *customerFromRepo.LastPurchaseStore, *customerHistory.LastPurchaseStore)
			}
		}
	})
}

var MockFileContent = `'CPF                PRIVATE     INCOMPLETO  DATA DA ÚLTIMA COMPRA TICKET MÉDIO          TICKET DA ÚLTIMA COMPRA LOJA MAIS FREQUÊNTE LOJA DA ÚLTIMA COMPRA
041.091.641-25     0           0           NULL                  NULL                  NULL                    NULL                NULL
058.189.421-97     0           0           NULL                  NULL                  NULL                    NULL                NULL
035.899.469-11     0           0           2011-12-07            972,40                972,40                  79.379.491/0008-50  79.379.491/0008-50
140.979.231-53     0           0           2011-12-07            308,06                308,06                  79.379.491/0008-51  79.379.491/0008-50
754.505.779-15     0           0           2011-12-07            999,00                999,00                  79.379.491/0008-50  79.379.491/0008-51
035.260.619-39     1           0           2011-12-07            1399,00               1399,00                 79.379.491/0008-50  79.379.491/0008-50`

var MockCustomerHistorySlicePersisted = []c.CustomerHistory{
	c.NewCustomerHistory("04109164125", true, ptrBool(false), ptrBool(false), nil, nil, nil, nil, nil, false, false),
	c.NewCustomerHistory("05818942197", false, ptrBool(false), ptrBool(false), nil, nil, nil, nil, nil, false, false),
	c.NewCustomerHistory("03589946911", true, ptrBool(false), ptrBool(false), ptrString("2011-12-07"), ptrFloat64(972.40), ptrFloat64(972.40), ptrString("79379491000850"), ptrString("79379491000850"), true, true),
	c.NewCustomerHistory("14097923153", true, ptrBool(false), ptrBool(false), ptrString("2011-12-07"), ptrFloat64(308.06), ptrFloat64(308.06), ptrString("79379491000851"), ptrString("79379491000850"), false, true),
	c.NewCustomerHistory("75450577915", true, ptrBool(false), ptrBool(false), ptrString("2011-12-07"), ptrFloat64(999.00), ptrFloat64(999.00), ptrString("79379491000850"), ptrString("79379491000851"), true, false),
	c.NewCustomerHistory("03526061939", true, ptrBool(true), ptrBool(false), ptrString("2011-12-07"), ptrFloat64(1399.00), ptrFloat64(1399.00), ptrString("79379491000850"), ptrString("79379491000850"), true, true),
}

func ptrBool(b bool) *bool {
	return &b
}

func ptrFloat64(f float64) *float64 {
	return &f
}

func ptrString(s string) *string {
	return &s
}
