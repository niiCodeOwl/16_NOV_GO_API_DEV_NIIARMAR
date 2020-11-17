// connection to db

package connection

import (
	  "gorm.io/driver/postgres"
	  "gorm.io/gorm"
	 
)

var DB *gorm.DB

func ConnectDataBase() {

  dsn := "user=postgres password= dbname=postgres port=5432 host=localhost sslmode=disable"
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Users{})
  DB = database

}