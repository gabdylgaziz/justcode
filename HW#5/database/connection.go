package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"bookstore/book"
	"sort"
	"os"
)


func GetBooks() []book.Book{
	var boooks []book.Book
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.AutoMigrate(&book.Book{})
	db.Find(&boooks)

	return boooks
}

func FilterByPriceAsc() []book.Book{
	items := GetBooks()
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price < items[j].Price
	})
	return items
}
