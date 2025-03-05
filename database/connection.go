package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// Connection establishes and returns a database connection
func Connection() (*sql.DB, error) {
	// Define MySQL connection parameters
	user := "chetanbudathoki"    // Change to your MySQL username
	password := "HeroBudathoki"  // Change to your MySQL password
	host := "94.136.185.141"     // Change to your MySQL server IP
	port := 9000                 // MySQL port
	database := "chetanbudathoki" // Change to your database name

	// Connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	// Open a connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("❌ Connection failed:", err)
		return nil, err
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Println("❌ Database is unreachable:", err)
		db.Close() // Close the DB before returning
		return nil, err
	}

	log.Println("✅ Connected to MySQL 8 successfully!")
	return db, nil
}
