package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
)

func TestUserDB(t *testing.T) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "admin",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "userdDB",
		AllowNativePasswords: true,
	}
	var err error
	usersDbtest, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		t.Fatalf("error opening database connection: %v", err)
	}

	pingErr := usersDbtest.Ping()
	if pingErr != nil {
		t.Fatalf("error pinging database: %v", pingErr)
	}

	var version string
	usersDbtest.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
func TestProdDB(t *testing.T) {
	productCachetest := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := productCachetest.Ping().Result()
	if err != nil {
		t.Fatalf("error pinging database: %v", err)
	}
	fmt.Printf("Connected to DB): %s", pong)
}
func TestKeywordDB(t *testing.T) {
	productCachetest := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	pong, err := productCachetest.Ping().Result()
	if err != nil {
		t.Fatalf("error pinging database: %v", err)
	}
	fmt.Printf("Connected to DB): %s", pong)
}

//	func TestGetProdAPI(t *testing.T) {
//		if getProdListFromApi("https://www.amazon.com/s?k=iphone") == nil {
//			t.Fatalf("Cant get Data from API")
//		}
//		fmt.Printf("Get Data from API Test Completed\n")
//	}
func TestGetBsrCrawler(t *testing.T) {
	if getBsr("https://www.amazon.com/dp/B07KJGKP9H") == nil {
		t.Fatalf("error getting bsr from crawler")
	}
	fmt.Printf("Get BSR from crawler Test Completed\n")
}
func TestWebPage(t *testing.T) {
	fmt.Printf("/ping request to the site test\n")
}
