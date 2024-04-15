package main

import (
	"bwa-go/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/bwago?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good")

	var users []user.User
	length := len(users)
	fmt.Println("cek panjang data : ", length)

	db.Find(&users)
	length = len(users)
	fmt.Println("cek panjang data : ", length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("++++++++++++++++++++++++++")
	}
}
