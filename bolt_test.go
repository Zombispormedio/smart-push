package main
import (
    "testing"
    "github.com/boltdb/bolt"
)



func TestMain(m *testing.T){
    
     db, err := bolt.Open("main.db", 0600, nil)
    if err != nil {
       m.Fatal(err);
    }
    defer db.Close()
    
    db.Update(func(tx *bolt.Tx) error{
         b, _ :=tx.CreateBucketIfNotExists([]byte("Config"))
      
         
        return b.Put([]byte("identifier"), []byte("n6NH976vNOHlWQwGH83uvXS9bTsrUtYb"))
    })
    
  m.Log("hello")
    
}