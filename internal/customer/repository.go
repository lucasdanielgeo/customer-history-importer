package customer

type CustomerHistoryRepository interface {
	SaveBatch(customers []CustomerHistory) error
}
