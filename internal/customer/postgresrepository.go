package customer

import (
	"database/sql"
	"fmt"
)

type CustomerHistoryDBRepository struct {
	db *sql.DB
}

func NewPostgresCustomerHistoryRepository(db *sql.DB) CustomerHistoryRepository {
	return CustomerHistoryDBRepository{db: db}
}

func (r CustomerHistoryDBRepository) SaveBatch(customers []CustomerHistory) error {
	query := `
        INSERT INTO customers_purchase_history (
            cpf, private, incomplete, last_purchase_date, average_ticket, 
            last_purchase_ticket, most_frequent_store, last_purchase_store,
			is_valid_cpf, is_most_frequent_store_valid, is_last_purchase_store_valid
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 ,$10 ,$11)
    `
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer stmt.Close()

	for _, c := range customers {
		_, err = stmt.Exec(
			c.CPF, c.Private, c.Incomplete, c.LastPurchaseDate, c.AverageTicket,
			c.LastPurchaseTicket, c.MostFrequentStore, c.LastPurchaseStore,
			c.IsValidCPF, c.IsValidMostFrequentStore, c.IsValidLastPurchaseStore,
		)
		if err != nil {
			return fmt.Errorf("error executing SQL statement: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing tx: %w", err)
	}

	return nil
}
