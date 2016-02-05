package db

import(
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func checkErr(err error, message string) {
  if err != nil{
    log.Fatal(message)
  }
}

func InitConn() (gorm.DB) {
  db, err := gorm.Open("postgres", "user=lastoller dbname=golinks_development sslmode=disable")
  checkErr(err, "DataBase connection failed!")
  return db
}
