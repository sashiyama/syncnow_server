package repository

import (
	"database/sql"
	"log"
)

type TransactionRepository struct {
	DB *sql.DB
}

func (tr *TransactionRepository) Transaction(txFunc func(*sql.Tx) (interface{}, error)) (interface{}, error) {
	log.Println("TRANSACTION [BEGIN]")
	tx, err := tr.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Println("TRANSACTION [ROLLBACK]")
			panic(p)
		} else if err != nil {
			tx.Rollback()
			log.Println("TRANSACTION [ROLLBACK]")
		} else {
			err = tx.Commit()
			log.Println("TRANSACTION [COMMIT]")
		}
	}()

	r, err := txFunc(tx)
	return r, err
}
