package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	username := flag.String("u", "", "Username for the new user")
	password := flag.String("p", "", "Password for the new user")
	salt := flag.String("s", "", "Salt for the password")
	dbPath := flag.String("d", "", "Path to the SQLite database")
	flag.Parse()

	if *username == "" || *password == "" || *salt == "" || *dbPath == "" {
		fmt.Println("Usage: -u USERNAME -p PASSWORD -s SALT -d DATABASE_PATH")
		os.Exit(1)
	}

	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password+*salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure the users table exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		userId INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Insert the new user
	insertUserQuery := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = db.Exec(insertUserQuery, *username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User registered successfully")
}
