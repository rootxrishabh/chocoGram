// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// )

// func GetDB() *gorm.DB {
// 	user := os.Getenv("MYSQL_USER")
// 	password := os.Getenv("MYSQL_PASSWORD")
// 	dbName := os.Getenv("MYSQL_DB")

// 	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbName)
// 	db, err := gorm.Open("mysql", dsn)
// 	if err != nil {
// 		log.Println("Error in connecting to the database")
// 	}
// 	return db
// }

package config

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/?charset=utf8mb4&parseTime=True&loc=Local", user, password)

	// Open a connection to MySQL without specifying a database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to MySQL")
	}

	// Create the database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic("Failed to create database")
	}

	// Close the connection
	db.Close()

	// Open a new connection to the newly created database
	dsn = fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbName)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return gormDB
}
