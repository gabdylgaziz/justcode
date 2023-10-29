package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/go-redis/redis/v8"
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Phone     string `db:"phone"`
	Login     string `db:"login"`
	Password  string `db:"password"`
}

var ctx = context.Background()

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
		DB:   0,               
	})
	defer redisClient.Close()

	db, err := sql.Open("postgres", "host=localhost user=postgres password=12345 dbname=payservice port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for{
		var key string
		fmt.Print("Write the key: ")
		fmt.Scan(&key)
		cachedData, err := redisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			var user User
			query := "SELECT * FROM users WHERE username = 'gabdyq'"
			rows, err := db.Query(query)
	
			for rows.Next() {
				err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Login, &user.Password)
				if err != nil {
					panic(err)
				}
			}
			redisClient.Set(ctx, key, user, time.Hour)
			fmt.Println("PSQL:", user)
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Redis:", cachedData)
		}
	}
	
}
