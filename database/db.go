package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
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

		prefix := []byte("bank_")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
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

func IsBankExists(name, code string) (bool, error) {
	exists := false
	err := DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte("bank_")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				var bank models.Bank
				if err := json.Unmarshal(val, &bank); err != nil {
					return err
				}
				if bank.Name == name || bank.Code == code {
					exists = true
					return nil
				}
				return nil
			})
			if err != nil {
				return err
			}
			if exists {
				return nil
			}
		}
		return nil
	})
	return exists, err
}

func AddBank(bank *models.Bank) error {
	exists, err := IsBankExists(bank.Name, bank.Code)
	if err != nil {
		return fmt.Errorf("error checking bank existence: %v", err)
	}
	if exists {
		return fmt.Errorf("bank with name '%s' or code '%s' already exists", bank.Name, bank.Code)
	}

	return DB.Update(func(txn *badger.Txn) error {
		if bank.ID == "" {
			bank.ID = uuid.New().String()
		}

		bankData, err := json.Marshal(bank)
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprintf("bank_%s", bank.ID))
		return txn.Set(key, bankData)
	})
}

func GetBankCount() (int, error) {
	count := 0
	err := DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte("bank_")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			count++
		}
		return nil
	})
	return count, err
}
