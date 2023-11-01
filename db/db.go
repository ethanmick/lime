package db

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func GetDB() (*bolt.DB, error) {
	db, err := bolt.Open("lime.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("emails"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db, nil
}
