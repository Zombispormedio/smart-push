package store

import (
	"github.com/boltdb/bolt"
)

func OpenDB() (*bolt.DB, error) {

	db, err := bolt.Open(".store/main.db", 0600, nil)
	return db, err

}

func Get(key string, bucket string, cb func(string)) error {

	var Error error

	db, Error := OpenDB()

	if Error == nil {

		Error = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))

			v := string(b.Get([]byte(key)))
			db.Close()
			cb(v)
			return nil
		})

	}

	return Error
}

func Put(key string, value string, bucket string) error {

	var Error error

	db, Error := OpenDB()

	if Error == nil {

		defer db.Close()

		Error = db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bucket))

			if err == nil {
				err = b.Put([]byte(key), []byte(value))

			}
			return err
		})

	}

	return Error
}

func PutWithDB(db *bolt.DB, key string, value string, bucket string) error {

	var Error error

	Error = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))

		if err == nil {
			err = b.Put([]byte(key), []byte(value))

		}
		return err
	})

	return Error
}

func GetWithDB(db *bolt.DB,  bucket string , key string) (string, error) {

	var Error error
	var Result string

	Error = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		Result = string(b.Get([]byte(key)))
		return nil
	})

	return Result, Error
}

func Iterate(db *bolt.DB, bucketName string, cb func(*bolt.Cursor) error) error {
	var Error error

	Error = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		return cb(c)
	})

	return Error
}

func DeleteWithDB(db *bolt.DB, key []byte,bucket string) error {

	var Error error

	Error = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))		
		return b.Delete(key)
	})

	return Error
}



