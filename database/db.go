package database

import (
	"encoding/json"
	"log"

	"github.com/dgraph-io/badger/v4"
	"github.com/nigeria-banks-api/models"
)

var DB *badger.DB

func InitDB() {
	var err error
	opts := badger.DefaultOptions("./badger-data")
	opts.Logger = nil
	DB, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func GetAllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	
	err := DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				var bank models.Bank
				if err := json.Unmarshal(val, &bank); err != nil {
					return err
				}
				banks = append(banks, bank)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return banks, nil
}

func AddBank(bank *models.Bank) error {
	return DB.Update(func(txn *badger.Txn) error {
		bankData, err := json.Marshal(bank)
		if err != nil {
			return err
		}
		
		key := []byte(bank.Code)
		return txn.Set(key, bankData)
	})
}

func GetBankCount() (int, error) {
	count := 0
	err := DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			count++
		}
		return nil
	})
	return count, err
}
