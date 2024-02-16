package customer

import (
	"fmt"
)

type MemoryCustomerHistoryRepository struct {
	Data map[string]CustomerHistory
}

func NewMemoryCustomerHistoryRepository() *MemoryCustomerHistoryRepository {
	return &MemoryCustomerHistoryRepository{
		Data: make(map[string]CustomerHistory),
	}
}

func (r *MemoryCustomerHistoryRepository) SaveBatch(customers []CustomerHistory) error {
	for _, c := range customers {
		r.Data[c.CPF] = c
	}
	return nil
}

func (r *MemoryCustomerHistoryRepository) Read(CPF string) (*CustomerHistory, error) {
	if customer, ok := r.Data[CPF]; ok {
		return &customer, nil
	}
	return nil, fmt.Errorf("customer with CPF %s not found", CPF)
}
