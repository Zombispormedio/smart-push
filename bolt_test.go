package test
import (
    "log"
    "testing"
    "github.com/boltdb/bolt"
)


func TestMain(m *testing.M){
     db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}