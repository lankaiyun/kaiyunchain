package db

import (
	"log"

	"github.com/cockroachdb/pebble"
)

func GetDbObj(path string) *pebble.DB {
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		log.Fatal("Failed to pebble.Open: ", err)
	}
	return db
}

func Set(k, v []byte, dbObj *pebble.DB) {
	if err := dbObj.Set(k, v, pebble.Sync); err != nil {
		log.Fatal("Failed to db.Set: ", err)
	}
}

func Get(k []byte, dbObj *pebble.DB) []byte {
	value, closer, err := dbObj.Get(k)
	if err != nil {
		return nil
	}
	if err := closer.Close(); err != nil {
		log.Panic("Failed to db.Close: ", err)
	}
	return value
}

func CloseDbObj(dbObj *pebble.DB) {
	if err := dbObj.Close(); err != nil {
		log.Panic("Failed to db.Close: ", err)
	}
}
