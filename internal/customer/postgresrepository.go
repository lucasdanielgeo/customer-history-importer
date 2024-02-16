package customer

import (
	"database/sql"
	"fmt"
	"log"
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

	var rowsInsertedCount int64
	for _, c := range customers {
		result, err := stmt.Exec(
			c.CPF, c.Private, c.Incomplete, c.LastPurchaseDate, c.AverageTicket,
			c.LastPurchaseTicket, c.MostFrequentStore, c.LastPurchaseStore,
			c.IsValidCPF, c.IsValidMostFrequentStore, c.IsValidLastPurchaseStore,
		)
		if err != nil {
			return fmt.Errorf("error executing SQL statement: %w", err)
		}

		rows, err := result.RowsAffected()
		if err != nil {
			log.Printf("[ERROR] could not get rows affected: %v", err)
		}
		rowsInsertedCount += rows
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing tx: %w", err)
	}

	log.Printf("[INFO] Sucessfuly inserted %d rows", rowsInsertedCount)

	return nil
}
