package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	// "gorm.io/driver/postgres"
	// "gorm.io/driver/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// gorm struct table
type Users struct {

	Id  	int64  `gorm:"primaryKey;autoIncrement;json:"id"`
	Hostname    string  `json:"hostname"`
	Address     string  `json:"address"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Port        string  `json:"port"`

}

var user []Users

//@route get all users
func getUsers(c echo.Context) error {
	users, _ := GetRepoUsers()
	return c.JSON(http.StatusOK, users)
}

func postUser(c echo.Context) error {
	// database connection
	db, err := gorm.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	u := new(Users)
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := "INSERT INTO users (hostname, address,username,password,port)VALUES ($1, $2, $3, $4, $5)"
	res, err := db.DB().Query(sqlStatement, u.Hostname, u.Address, u.Username, u.Password, u.Port)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}




  

var DB *gorm.DB

func GetRepoUsers() ([]Users, error) {

	// database connection
	db, err := gorm.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	users := []Users{}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}



// main function of the pg
func main() {

	e := echo.New()
	e.GET("/data",getUsers)
	e.POST("/echo",postUser)
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	

	// database connection
	db, err := gorm.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&Users{})
	db.CreateTable(&Users{})
	
	// panic("Connected")
	e.Logger.Fatal(e.Start(":8080"))

}



